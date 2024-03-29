package consensus

import (
	"github.com/adithyabhatkajake/libchatter/crypto"
	"github.com/adithyabhatkajake/libchatter/log"
	"github.com/adithyabhatkajake/libsynchs/chain"
	msg "github.com/adithyabhatkajake/libsynchs/msg"
)

// How to create and validate certificates

// NewCert creates a certificate
func NewCert(certMap map[uint64]*msg.Vote, blockhash crypto.Hash, view uint64) *msg.BlockCertificate {
	bc := &msg.BlockCertificate{}
	bc.SetBlockInfo(blockhash, view)
	bc.Init()
	for _, v := range certMap {
		bc.AddVote(*v)
	}
	return bc
}

// IsCertValid checks if the certificate is valid for the data
func (n *SyncHS) IsCertValid(bc *msg.BlockCertificate) bool {
	log.Debug("Received a block certificate -")
	// Certificate for genesis is always correct
	if h, _ := bc.GetBlockInfo(); h == chain.EmptyHash {
		return true
	}
	if bc.GetNumSigners() <= n.GetNumberOfFaultyNodes() {
		log.Error("The certificate has <= f signatures")
		return false
	}
	for idx, id := range bc.GetSigners() {
		sig := bc.GetSignatureFromID(id)
		if sig == nil {
			log.Error("Signature for ID not found")
			return false
		}
		sigOk, err := n.GetPubKeyFromID(id).Verify(bc.GetData(), sig)
		if err != nil {
			log.Error("Certificate signature verification error")
			return false
		}
		if !sigOk {
			log.Error("Certificate signature is invalid for idx", idx)
			return false
		}
	}
	return true
}

func (n *SyncHS) addCert(bc *msg.BlockCertificate, blockNum uint64) {
	log.Debug("Adding certificate to block ", blockNum)
	n.certMapLock.Lock()
	n.certMap[blockNum] = bc
	n.certMapLock.Unlock()
}

func (n *SyncHS) getCertForBlockIndex(idx uint64) (*msg.BlockCertificate, bool) {
	n.certMapLock.Lock()
	defer n.certMapLock.Unlock()
	blk, exists := n.certMap[idx]
	return blk, exists
}
