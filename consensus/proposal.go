package consensus

import (
	"time"

	"github.com/adithyabhatkajake/libchatter/crypto"
	"github.com/adithyabhatkajake/libchatter/log"
	"github.com/adithyabhatkajake/libchatter/util"
	"github.com/adithyabhatkajake/libsynchs/chain"
	msg "github.com/adithyabhatkajake/libsynchs/msg"
	pb "github.com/golang/protobuf/proto"
)

// Monitor pending commands, if there is any change and the current node is the leader, then start proposing blocks
func (n *SyncHS) propose() {
	log.Debug("Starting a propose step")
	// Do we have a certificate?
	n.bc.Mu.Lock()
	defer n.bc.Mu.Unlock()
	head := n.bc.Head
	cert, exists := n.getCertForBlockIndex(head)
	if !exists {
		log.Debug("The head does not have a certificate")
		log.Debug("Cancelling the proposal")
		return
	}
	cmds, isSuff := n.getCmdsIfSufficient()
	if !isSuff {
		log.Debug("Insufficient commands, aborting the proposal")
		return
	}
	n.bc.Head++
	newHeight := n.bc.Head
	prop := n.NewCandidateProposal(cmds, cert, newHeight, nil)
	block := &chain.ExtBlock{}
	block.FromProto(prop.Block)
	// Add this block to the chain, for future proposals
	n.bc.BlocksByHeight[newHeight] = block
	n.bc.BlocksByHash[block.GetBlockHash()] = block
	log.Trace("Finished Proposing")
	// Ship proposal to processing
	relayMsg := &msg.SyncHSMsg{}
	relayMsg.Msg = &msg.SyncHSMsg_Prop{Prop: prop}
	log.Debug("Proposing block:", prop.String())
	go func() {
		// Leader sends new block to all the other nodes
		n.Broadcast(relayMsg)
		// Leader should also vote
		n.voteForBlock(block)
		// Start 2\delta timer
		n.startBlockTimer(block)
	}()
}

// Deal with the proposal
func (n *SyncHS) proposeHandler(prop *msg.Proposal) {
	if prop.Miner != n.leader {
		log.Info("Proposal received from invalid node, not leader")
		return
	}
	ht := prop.Block.GetHeader().GetHeight()
	log.Trace("Handling proposal ", ht)
	ep := &msg.ExtProposal{}
	ep.FromProto(prop)
	if crypto.ToHash(ep.Block.BlockHash) != ep.GetBlockHash() {
		log.Warn("Invalid block. Computed Hash and the Obtained hash does not match")
		return
	}
	data, _ := pb.Marshal(prop.GetBlock().GetHeader())
	correct, err := n.GetPubKeyFromID(n.leader).Verify(data, ep.GetMiningProof())
	if !correct || err != nil {
		log.Error("Incorrect signature for proposal ", ht)
		return
	}
	// Check block certificate for non-genesis blocks
	if !n.IsCertValid(&ep.BlockCertificate) {
		log.Error("Invalid certificate received for block", ht)
		return
	}
	var blk *chain.ExtBlock
	var exists bool
	{
		// First check for equivocation
		n.bc.Mu.RLock()
		blk, exists = n.bc.BlocksByHeight[ht]
		n.bc.Mu.RUnlock()
	}
	if exists && ep.GetBlockHash() != blk.GetBlockHash() {
		// Equivocation
		log.Warn("Equivocation detected.", ep.GetBlockHash(), blk.GetBlockHash())
		// TODO trigger view change
		return
	}
	if exists {
		// Duplicate block received,
		// we have already committed this block, IGNORE
		return
	}
	n.addNewBlock(&ep.ExtBlock)
	n.ensureBlockIsDelivered(&ep.ExtBlock)

	// Vote for the proposal
	go func() {
		n.voteForBlock(&ep.ExtBlock)
		// Start 2\delta timer
		n.startBlockTimer(&ep.ExtBlock)
		// Stop blame timer, since we got a valid proposal
		// During commit, if pending commands is empty, we will restart the blame timer
		n.stopBlameTimer()
	}()

}

// NewBlockBody creates a new block body from the commands received.
func NewBlockBody(cmds [][]byte) *chain.ExtBody {
	bd := &chain.ExtBody{}
	bd.Txs = cmds
	return bd
}

func (n *SyncHS) ensureBlockIsDelivered(blk *chain.ExtBlock) {
	var exists bool
	var parentblk *chain.ExtBlock
	// Ensure that all the parents are delivered first.
	parentIdx := blk.GetHeight() - 1
	// Wait for parents to be delivered first
	for tries := 30; tries > 0; tries-- {
		<-time.After(time.Millisecond)
		n.bc.Mu.RLock()
		parentblk, exists = n.bc.BlocksByHeight[parentIdx]
		n.bc.Mu.RUnlock()
		if exists && parentblk.GetBlockHash() != blk.GetParentHash() {
			// This block is delivered.
			log.Warn("Block  ", blk.GetHeight(), " extending wrong parent.\n",
				"Wanted Parent Block:", util.HashToString(parentblk.GetBlockHash()),
				"Found Parent Block:", util.HashToString(blk.GetParentHash()))
			return
		}
		if exists {
			// The parent of the proposed block is the same as the block we have at the parent's position, CONTINUE
			break
		}
	}
	if !exists {
		// The parents are not delivered, so we cant process this block
		// Return
		log.Warn("Parents not delivered, aborting this proposal")
		return
	}
	// All parents are delivered, lets break out!!
	log.Trace("All parents are delivered")
}

func (n *SyncHS) startBlockTimer(blk *chain.ExtBlock) {
	// Start 2delta timer
	timer := util.NewTimer(func() {
		log.Info("Committing block-", blk.GetHeight())
		// We have committed this block
		// Let the client know that we committed this block
		synchsmsg := &msg.SyncHSMsg{}
		ack := &msg.SyncHSMsg_Ack{}
		ack.Ack = &msg.CommitAck{
			Block: blk.ToProto(),
		}
		synchsmsg.Msg = ack
		// Tell all the clients, that I have committed this block
		n.ClientBroadcast(synchsmsg)
		// }
	})
	log.Info("Started timer for block-", blk.GetHeight())
	timer.SetTime(n.GetCommitWaitTime())
	n.addNewTimer(blk.GetHeight(), timer)
	timer.Start()
}

func (n *SyncHS) addNewBlock(blk *chain.ExtBlock) {
	// Otherwise, add the current block to map
	n.bc.Mu.Lock()
	n.bc.BlocksByHeight[blk.GetHeight()] = blk
	n.bc.BlocksByHash[blk.GetBlockHash()] = blk
	n.bc.Mu.Unlock()
}

func (n *SyncHS) addNewTimer(pos uint64, timer *util.Timer) {
	n.timerLock.Lock()
	n.timerMaps[pos] = timer
	n.timerLock.Unlock()
}

// NewCandidateProposal returns a proposal message built using commands
func (n *SyncHS) NewCandidateProposal(cmds [][]byte,
	cert *msg.BlockCertificate, newHeight uint64, extra []byte) *msg.Proposal {
	bhash, view := cert.GetBlockInfo()
	// Start setting block fields
	pbody := &chain.ProtoBody{
		Txs:       cmds,
		Responses: cmds, // For now, the response is the same as the cmd
	}
	pheader := &chain.ProtoHeader{
		Extra:      extra,
		Height:     newHeight,
		ParentHash: bhash.GetBytes(),
		TxHash:     nil, // Compute merkle tree out of transactions in the block body
	}
	// Set Hash
	log.Debug("PrevHash:",
		util.HashToString(crypto.ToHash(pheader.GetParentHash())))
	log.Debug("Computed Proposal ", newHeight,
		" with hash ", util.HashToString(bhash))
	// Sign
	data, _ := pb.Marshal(pheader)
	newBlockHash := crypto.DoHash(data)
	sig, err := n.GetMyKey().Sign(data)
	if err != nil {
		log.Error("Error in signing a block during proposal")
		panic(err)
	}
	blk := &chain.ProtoBlock{
		Header:    pheader,
		Body:      pbody,
		BlockHash: newBlockHash.GetBytes(),
	}
	// Build Propose Evidence
	pevidence, _ := pb.Marshal(cert.ToProto())
	prop := &msg.Proposal{
		Miner:           n.GetId(),
		View:            view,
		Block:           blk,
		MiningProof:     sig,       // Signature from the leader in the current view
		ProposeEvidence: pevidence, // Certificate for parent block
	}
	return prop
}
