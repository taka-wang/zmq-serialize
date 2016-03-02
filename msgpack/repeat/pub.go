package main

import (
	"bytes"
	_ "encoding/json"
	_ "fmt"
	zmq "github.com/taka-wang/zmq3"
	"github.com/ugorji/go/codec"
	"time"
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
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Id   int    `json:"id"`
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
	pub()
}

func pub() {
	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Connect("ipc:///tmp/dummy")

	command := MbTcpMultipleWriteReq{ // named key
		CmdHeader: CmdHeader{
			Receiver: "core",
			Sender:   "me",
			Version:  "1",
			Tid:      32,
			Method:   "mbtcp.once.write",
		},
		MbTcpHeader: MbTcpHeader{
			Ip:   "192.168.1.1",
			Port: 503,
			Id:   22,
		},
	}

	command.Requests = append(command.Requests,
		MbWriteRequest{
			Code:     1,
			Register: 2003,
			Value:    "1025",
			Type:     "int64",
			Alias:    "hello",
		},
	)
	command.Requests = append(command.Requests,
		MbWriteRequest{
			Code:     1,
			Register: 2003,
			Value:    "1025",
			Type:     "int64",
			Alias:    "hello",
		},
	)

	// pack
	buf := &bytes.Buffer{}
	enc := codec.NewEncoder(buf, mh)
	enc.Encode(command)

	for {
		sender.Send("mbtcp.once.write", zmq.SNDMORE)
		sender.Send(buf.String(), 0) // convert buffer arr to str
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
