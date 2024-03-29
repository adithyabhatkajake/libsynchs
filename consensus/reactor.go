package consensus

import (
	"github.com/adithyabhatkajake/libchatter/log"
	msg "github.com/adithyabhatkajake/libsynchs/msg"
	pb "github.com/golang/protobuf/proto"
)

func (n *SyncHS) react(m []byte) {
	log.Trace("Received a message of size", len(m))
	inMessage := &msg.SyncHSMsg{}
	err := pb.Unmarshal(m, inMessage)
	if err != nil {
		log.Error("Received an invalid protocol message from client", err)
		return
	}
	n.msgChannel <- inMessage
}

func (n *SyncHS) protocol() {
	// Process protocol messages
	for {
		msgIn, ok := <-n.msgChannel
		if !ok {
			log.Error("Msg channel error")
			return
		}
		log.Trace("Received msg", msgIn.String())
		switch x := msgIn.Msg.(type) {
		case *msg.SyncHSMsg_Prop:
			prop := msgIn.GetProp()
			log.Debug("Received a proposal from ", prop.GetMiner())
			// Send proposal to propose handler
			go n.proposeHandler(prop)
		case *msg.SyncHSMsg_Npblame:
			blMsg := msgIn.GetNpblame()
			go n.handleNoProgressBlame(blMsg)
		case *msg.SyncHSMsg_Eqblame:
			_ = msgIn.GetEqblame()
			// TODO
		case *msg.SyncHSMsg_Vote:
			pvote := msgIn.GetVote()
			vote := &msg.Vote{}
			vote.FromProto(pvote)
			go func() {
				n.voteChannel <- vote
			}()
		case nil:
			log.Warn("Unspecified msg type", x)
		default:
			log.Warn("Unknown msg type", x)
		}
	}
}
