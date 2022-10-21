package consensus

import (
	"bufio"
	"context"
	"sync"

	"github.com/adithyabhatkajake/libchatter/util"

	chain "github.com/adithyabhatkajake/libsynchs/chain"
	config "github.com/adithyabhatkajake/libsynchs/config"
	msg "github.com/adithyabhatkajake/libsynchs/msg"

	"github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
)

// SyncHS implements the consensus protocol!!
type SyncHS struct {
	// Network data structures
	host    host.Host
	cliHost host.Host
	ctx     context.Context

	// Maps
	// Mapping between ID and libp2p-peer
	pMap map[uint64]*peerstore.AddrInfo
	// A set of all known clients
	cliMap map[*bufio.ReadWriter]bool
	// A map of node ID to its corresponding RW stream
	streamMap map[uint64]*bufio.ReadWriter
	// A map of hash to pending commands
	pendingCommands [][]byte
	// A mapping between the block number to its commit timer
	timerMaps map[uint64]*util.Timer
	// Certificate map
	certMap map[uint64]*msg.BlockCertificate
	// A mapping between the view and (A mapping between the origin and blames against the leader)
	// blameMap map[uint64]map[uint64]*msg.Blame
	// correct vote map (recorder:origin:the number/value of vote/proposal/reputation)
	voteMap map[uint64]map[uint64]map[uint64]uint64
	// malicous vote map
	voteMaliMap map[uint64]map[uint64]map[uint64]uint64
	// correct proposal map
	proposalMap map[uint64]map[uint64]map[uint64]uint64
	//equivocate proposal map
	equiproposalMap map[uint64]map[uint64]map[uint64]uint64
	//withholding proposal map
	withproposalMap map[uint64]map[uint64]map[uint64]uint64
	//malicious proposal map
	maliproposalMap map[uint64]map[uint64]map[uint64]uint64
	//Reputation map
	reputationMap map[uint64]map[uint64]map[uint64]uint64
	//ProosalByheightMap
	proposalByviewMap map[uint64]*msg.Proposal
	/* Locks - We separate all the locks, so that acquiring
	one lock does not make other goroutines stop */
	cliMutex           sync.RWMutex // The lock to modify cliMap
	netMutex           sync.RWMutex // The lock to modify streamMap: Use mutex when using network streams to talk to other nodes
	cmdMutex           sync.RWMutex // The lock to modify pendingCommands
	timerLock          sync.RWMutex // The lock to modify timerMaps
	certMapLock        sync.RWMutex // The lock to modify certMap
	repMapLock         sync.RWMutex // The lock to modify reputationMap
	voteMapLock        sync.RWMutex // The lock to modify reputationMap
	propMapLock        sync.RWMutex // The lock to modify reputationMap
	malipropLock       sync.RWMutex //........
	equipropLock       sync.RWMutex
	voteMaliLock       sync.RWMutex
	withpropoLock      sync.RWMutex
	proposalByviewLock sync.RWMutex

	// Channels
	msgChannel     chan *msg.SyncHSMsg // All messages come here first
	cmdChannel     chan []byte         // All commands are re-directed here
	voteChannel    chan *msg.Vote      // All votes are sent here
	proposeChannel chan *msg.Proposal  // All proposals are sent here
	errCh          chan error          // All errors are sent here

	// Block chain
	bc *chain.BlockChain

	// Protocol information
	leader uint64
	view   uint64

	// Embed the config
	*config.NodeConfig
}
