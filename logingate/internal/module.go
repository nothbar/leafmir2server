package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"leafmir2server/conf"
	"leafmir2server/login"
	"leafmir2server/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         fmt.Sprintf("%s:%d", conf.Server.TcpAddr, conf.Server.LoginTCPPort),
		Processor:       msg.Processor,
		AgentChanRPC:    login.ChanRPC,
		MsgParser:       NewMsgParser(),
	}
}
