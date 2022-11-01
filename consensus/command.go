package consensus

import "time"

func (n *SyncHS) addCmdsAndStartTimerIfSufficientCommands(cmd []byte) {
	n.cmdMutex.Lock()
	defer n.cmdMutex.Unlock()
	n.pendingCommands = append(n.pendingCommands, cmd)
	// .pendingCommands.PushBack(cmd)
	//&& n.GetID() == n.leader
	if uint64(len(n.pendingCommands)) >= n.GetBlockSize() { //Sufficient Commands start our timer!
		// change the condition of this: 1. if this block is gensis block start timer directly
		// 2. else
		if n.gcallFuncFinish {
			go n.startConsensusTimer()
			time.Sleep(time.Millisecond * 50)
			n.gcallFuncFinish = false
		}
		if !n.callFuncNotFinish {
			go n.startConsensusTimer()
			time.Sleep(time.Millisecond * 50)
			n.callFuncNotFinish = true
		}
		// go n.startConsensusTimer()
		// if uint64(len(n.pendingCommands)) >= n.GetBlockSize() {
		// 	go n.gnerateCandidateBlock()
		// }

		// }
	}
}

func (n *SyncHS) getCmdsIfSufficient() ([][]byte, bool) {
	blkSize := n.GetBlockSize()
	n.cmdMutex.Lock()
	defer n.cmdMutex.Unlock()
	numCmds := uint64(len(n.pendingCommands))
	if numCmds < blkSize {
		return nil, false
	}
	cmds := make([][]byte, blkSize)
	// Copy slice blkSize commands from pending Commands
	copy(cmds, n.pendingCommands[numCmds-blkSize:])
	// Update old slice
	n.pendingCommands = n.pendingCommands[:numCmds-blkSize]
	return cmds, true
}
