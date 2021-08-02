package internal

import "github.com/name5566/leaf/agent"

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(agent.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(agent.Agent)
	_ = a
}
