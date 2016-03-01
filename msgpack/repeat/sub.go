package main

import (
	_ "encoding/json"
	"fmt"
	zmq "github.com/taka-wang/zmq3"
	"github.com/ugorji/go/codec"
	_ "time"
)

var (
	mh = &codec.MsgpackHandle{RawToString: true}
)

type CmdHeader struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Version  string `json:"version"`
	Tid      int    `json:"tid"`
	Method   string `json:"method"`
}

type MbTcpHeader struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
	ID   int    `json:"id"`
}

type MbWriteRequest struct {
	Code     int    `json:"code"`
	Register int    `json:"register"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	Alias    string `json:"alias"`
}

type MbTcpSingleWriteReq struct {
	CmdHeader
	MbTcpHeader
	MbWriteRequest
}

type MbTcpMultipleWriteReq struct {
	CmdHeader
	MbTcpHeader
	Requests []MbWriteRequest `json:"requests"`
}

func main() {
	sub()
}

func sub() {
	receiver, _ := zmq.NewSocket(zmq.SUB)
	defer receiver.Close()
	receiver.Bind("ipc:///tmp/dummy")

	filter := "mbtcp"
	receiver.SetSubscribe(filter) // filter frame 1

	for {
		msg, _ := receiver.RecvMessage(0) // frame len: 2

		// unpack
		var s MbTcpMultipleWriteReq
		dec := codec.NewDecoderBytes([]byte(msg[1]), mh)
		err := dec.Decode(&s)

		if err != nil {
			fmt.Println(err) // unmarshal from json string failed
		} else {
			fmt.Println(msg[0])
			fmt.Println(s)
		}
	}
}
