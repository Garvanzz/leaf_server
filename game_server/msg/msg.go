package msg

import (
	"github.com/leaf/network"
)

const (
	HeadLenSize   = 2 //头部长度大小
	ProtoIdSize   = 4 //ProtoID大小
	ErrorCodeSize = 2 //错误码大小
	Key = "" //加密pb用
)

var Processor network.Processor

func init() {
	// init processor
}


//type Processor interface {
	//Route(msg interface{}, userData interface{}) error
	//Unmarshal(data []byte) (interface{}, error)
	//Marshal(msg interface{}) ([][]byte, error)
//}

type BProcessor struct {

}

func (bp *BProcessor)Register(){

}

func (bp *BProcessor)Route(msg interface{}, userData interface{}) error{

}

func (bp *BProcessor)Unmarshal(data []byte) (interface{}, error){

}

func (bp *BProcessor)Marshal(msg interface{}) ([][]byte, error){

}


