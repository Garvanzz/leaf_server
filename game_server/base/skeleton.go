package base

import (
	"game_server/conf"
	"github.com/leaf/chanrpc"
	"github.com/leaf/module"
)

func NewSkeleton()*module.Skeleton{
	skeleton:=&module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
