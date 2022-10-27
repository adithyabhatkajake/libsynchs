package consensus

func (n *SyncHS) addCmdsAndStartTimerIfSufficientCommands(cmd []byte) {
	n.cmdMutex.Lock()
	defer n.cmdMutex.Unlock()
	n.pendingCommands = append(n.pendingCommands, cmd)
	// n.pendingCommands.PushBack(cmd)
	if uint64(len(n.pendingCommands)) >= n.GetBlockSize() { //Sufficient Commands start our timer！
		go n.startConsensusTimer()
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
