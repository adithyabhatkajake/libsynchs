// this file finish how to compute node's reputation
package consensus

import (
	"math"
	"math/big"
	"sync"

	"github.com/adithyabhatkajake/libchatter/log"
)

const (
	pEpsilonWith     = float64(1.5)
	pEpsilonEqui     = float64(1.5)
	pEpsilonMali     = float64(1)
	vEpisilonMali    = float64(1)
	gamma            = float64(0.01)
	initialNodescore = float64(1e-6)
)

var (
	proposalnum     uint64
	votenum         uint64
	maliproposalnum uint64
	equiprospoalnum uint64
	withpropsoalnum uint64
	malivotenum     uint64
)

// TODO change correspongding reputaion map
func (n *SyncHS) ReputationCalculateinCurrentRound(nodeID uint64) *big.Float {
	//first we get the correct proposal/vote from map
	//get current various proposal/vote number
	// roundnum := n.view
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		n.proposalNumCalculate(nodeID)
		n.voteNumCalculate(nodeID)
		// n.maliproposalNumCalculate(nodeID)
		// n.equivocationproposalNumCalculate(nodeID)
		// n.withholdproposalNumCalculate(nodeID)
		// n.malivoteNumCalculate(nodeID)

	}()
	wg.Wait()
	// log.Info("calculate reputation for node", nodeID)
	proposalsc := new(big.Float).SetUint64(proposalnum)
	// - (float64(maliproposalnum)*pEpsilonMali +
	// float64(equiprospoalnum)*pEpsilonEqui +
	// float64(withpropsoalnum)*pEpsilonWith)
	proposalscore := n.maxValueCheckNum(proposalsc)
	votesc := new(big.Float).SetUint64(votenum)
	// - float64(malivotenum)*vEpisilonMali
	votescore := n.maxValueCheckNum(votesc)
	calInitialNodescore := new(big.Float).SetFloat64(initialNodescore)
	calGama := new(big.Float).SetFloat64(gamma)
	transcriptNum := new(big.Float).Add(votescore, proposalscore)
	gamaMulTranscript := new(big.Float).Mul(calGama, transcriptNum)
	fltnum, _ := gamaMulTranscript.Float64()
	behaviourSocre := new(big.Float).SetFloat64(math.Tanh(fltnum))
	baseNodeSocre := new(big.Float).Add(behaviourSocre, calInitialNodescore)
	nodeScore := n.maxValueCheckScore(baseNodeSocre)
	log.Info("node", n.GetID(), "calculate the reputation of", nodeID, "is", nodeScore)
	return nodeScore

}

func (n *SyncHS) reputationCountforRound() *big.Float {
	var currentReputationSum *big.Float = new(big.Float)
	for i := 0; i < len(n.pMap); i++ {
		currentReputationSum = currentReputationSum.Add(currentReputationSum, n.ReputationCalculateinCurrentRound(uint64(i)))
	}
	return currentReputationSum
}

func (n *SyncHS) proposalNumCalculate(nodeID uint64) uint64 {
	n.propMapLock.RLock()
	defer n.propMapLock.RUnlock()
	// _, exists := n.proposalMap[n.GetID()]
	// if exists {
	// 	for _, senderMap := range n.proposalMap[n.GetID()] {
	// 		num, exists1 := senderMap[nodeID]
	// 		if exists1 && num == 1 {
	// 			log.Debug("a valid num have been recored")
	// 			proposalnum++
	// 		}
	// 		// else {
	// 		// 	log.Debug(nodeID, "don't propose in this view")
	// 		// }
	// 	}
	// 	return proposalnum
	// }
	// return 0
	proposalnum = 0
	for _, senderMap := range n.proposalMap {
		num, exists := senderMap[nodeID]
		if exists && num == 1 {
			proposalnum++
		}

	}
	return proposalnum

}

func (n *SyncHS) voteNumCalculate(nodeID uint64) uint64 {
	n.voteMapLock.RLock()
	defer n.voteMapLock.RUnlock()
	// _, exists := n.voteMap[n.GetID()]
	// if exists {
	// 	for _, senderMap := range n.voteMap[n.GetID()] {
	// 		num, exists1 := senderMap[nodeID]
	// 		if exists1 && num == 1 {
	// 			votenum++
	// 		}
	// 		// else {
	// 		// 	log.Debug(nodeID, "don't vote in this view")
	// 		// }
	// 	}
	// 	return votenum
	// }

	// return 0
	votenum = 0
	for _, votermap := range n.voteMap {
		num, exists := votermap[nodeID]
		if exists && num == 1 {
			votenum++
		}

	}
	return votenum
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
	_, exists := n.voteMaliMap[n.GetID()]
	if exists {
		for _, senderMap := range n.voteMaliMap[n.GetID()] {

			num, exists1 := senderMap[nodeID]
			if exists1 && num == 1 {
				malivotenum++
			}

		}
		return malivotenum

	} else {
		return 0
	}

}

func (n *SyncHS) maxValueCheckNum(a *big.Float) *big.Float {
	// if a >= initialNodescore {
	// 	return a
	// } else {
	// 	return initialNodescore
	// }
	b := new(big.Float).SetUint64(0)
	c := a.Cmp(b)
	if c == -1 {
		return b
	} else {
		return a
	}

}

func (n *SyncHS) maxValueCheckScore(a *big.Float) *big.Float {
	b := new(big.Float).SetFloat64(initialNodescore)
	c := a.Cmp(b)
	if c == -1 {
		return b
	} else {
		return a
	}
}

//calcute the score of each node in this round

func (n *SyncHS) addNewViewReputaiontoMap() {
	n.repMapLock.Lock()
	defer n.repMapLock.Unlock()
	for i := uint64(0); i <= uint64(len(n.pMap)); i++ {
		// n.reputationMapwithoutRound[i] = n.ReputationCalculateinCurrentRound(i)
		if _, exists1 := n.reputationMap[n.view+1]; exists1 {
			n.reputationMap[n.view+1][i] = n.ReputationCalculateinCurrentRound(i)
		} else {
			n.reputationMap[n.view+1] = make(map[uint64]*big.Float)
			n.reputationMap[n.view+1][i] = n.ReputationCalculateinCurrentRound(i)
		}
	}
	log.Debug("Node", n.GetID(), "map is ", n.reputationMap)
}
