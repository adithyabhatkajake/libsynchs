// this file finish how to compute node's reputation
package consensus

import (
	"math"
	"sync"

	"github.com/adithyabhatkajake/libchatter/log"
)

const (
	pEpsilonWith  = 1.5
	pEpsilonEqui  = 1.5
	pEpsilonMali  = 1
	vEpisilonMali = 1
	gama          = 0.5
)

var (
	proposalnum     uint64 = 0
	votenum         uint64 = 0
	maliproposalnum uint64 = 0
	equiprospoalnum uint64 = 0
	withpropsoalnum uint64 = 0
	malivotenum     uint64 = 0
)

func (n *SyncHS) ReputationCalculateinCurrentRound(nodeID uint64) {
	//first we get the correct proposal/vote from map
	//get current various proposal/vote number
	// roundnum := n.view
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		n.proposalNumCalculate(nodeID)
		n.voteNumCalculate(nodeID)
		n.maliproposalNumCalculate(nodeID)
		n.equivocationproposalNumCalculate(nodeID)
		n.withholdproposalNumCalculate(nodeID)
		n.malivoteNumCalculate(nodeID)

	}()
	wg.Wait()
	log.Info("calculate reputation for node", nodeID)
	proposalsc := float64(proposalnum) - (float64(maliproposalnum)*pEpsilonMali +
		float64(equiprospoalnum)*pEpsilonEqui +
		float64(withpropsoalnum)*pEpsilonWith)
	proposalscore := n.maxvaluecheck(proposalsc)
	votesc := float64(votenum) - float64(malivotenum)*vEpisilonMali
	votescore := n.maxvaluecheck(votesc)
	nodeScore := math.Tanh(gama * (votescore + proposalscore))
	log.Info("The reputation of", nodeID, "is", nodeScore)

}

func (n *SyncHS) proposalNumCalculate(nodeID uint64) uint64 {
	n.propMapLock.RLock()
	defer n.propMapLock.RUnlock()
	_, exists := n.proposalMap[n.GetID()]
	if exists {
		for _, senderMap := range n.proposalMap[n.GetID()] {
			num, exists1 := senderMap[nodeID]
			if exists1 && num == 1 {
				proposalnum++
			} else {
				return 0
			}

		}
		return proposalnum
	} else {
		return 0
	}

}
func (n *SyncHS) voteNumCalculate(nodeID uint64) uint64 {
	n.voteMapLock.RLock()
	defer n.voteMapLock.RUnlock()
	_, exists := n.voteMap[n.GetID()]
	if exists {
		for _, senderMap := range n.voteMap[n.GetID()] {
			num, exists1 := senderMap[nodeID]
			if exists1 && num == 1 {
				votenum++
			} else {
				return 0
			}
		}
		return proposalnum
	} else {
		return 0
	}

}

func (n *SyncHS) maliproposalNumCalculate(nodeID uint64) uint64 {
	n.malipropLock.RLock()
	defer n.malipropLock.RUnlock()
	_, exists := n.maliproposalMap[n.GetID()]
	if exists {
		for _, senderMap := range n.maliproposalMap[n.GetID()] {
			num, exists1 := senderMap[nodeID]
			if exists1 && num == 1 {
				maliproposalnum++
			} else {
				return 0
			}
		}
		return maliproposalnum
	} else {
		return 0
	}

}

func (n *SyncHS) withholdproposalNumCalculate(nodeID uint64) uint64 {
	n.withpropoLock.RLock()
	defer n.withpropoLock.RUnlock()
	_, exists := n.withproposalMap[n.GetID()]
	if exists {
		for _, senderMap := range n.withproposalMap[n.GetID()] {
			num, exists1 := senderMap[nodeID]
			if exists1 && num == 1 {
				withpropsoalnum++
			} else {
				return 0
			}
		}
		return withpropsoalnum
	} else {
		return 0
	}

}

func (n *SyncHS) equivocationproposalNumCalculate(nodeID uint64) uint64 {
	n.equipropLock.RLock()
	defer n.equipropLock.RUnlock()
	_, exists := n.equiproposalMap[n.GetID()]
	if exists {
		for _, senderMap := range n.equiproposalMap[n.GetID()] {

			num, exists1 := senderMap[nodeID]
			if exists1 && num == 1 {
				equiprospoalnum++
			} else {
				return 0
			}

		}
		return equiprospoalnum

	} else {
		return 0
	}

}

func (n *SyncHS) malivoteNumCalculate(nodeID uint64) uint64 {
	n.voteMaliLock.RLock()
	defer n.voteMaliLock.RUnlock()
	_, exists := n.maliproposalMap[n.GetID()]
	if exists {
		for _, senderMap := range n.voteMaliMap[n.GetID()] {

			num, exists1 := senderMap[nodeID]
			if exists1 && num == 1 {
				malivotenum++
			} else {
				return 0
			}

		}
		return malivotenum

	} else {
		return 0
	}

}

func (n *SyncHS) maxvaluecheck(a float64) float64 {
	if a >= 0 {
		return a
	} else {
		return 0
	}
}

//calcute the score of each node in this round
