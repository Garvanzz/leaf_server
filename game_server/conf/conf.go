package conf

import (
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags

	// gate conf 网关配置
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 65535 //最大消息长度 原为4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2 //消息长度为2=uint16 不过也没用 内部重构大端
	LittleEndian           = false //这里设置无用 因为内部重构了默认大端

	// skeleton conf 框架配置
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
)
