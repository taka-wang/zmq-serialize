package main

import (
	"bytes"
	zmq "github.com/taka-wang/zmq3"
	"github.com/ugorji/go/codec"
	"time"
)

var (
	mh = &codec.MsgpackHandle{RawToString: true}

	command = MbTcpSingleWriteReq{ // named key
		CmdHeader: CmdHeader{
			Receiver: "mbtcp",
			Sender:   "restful",
			Version:  "1",
			Tid:      33,
			Method:   "mbtcp.once.write",
		},
		MbTcpHeader: MbTcpHeader{
			Ip:   "192.168.1.1",
			Port: 503,
			Id:   22,
		},
		MbWriteRequest: MbWriteRequest{
			Code:     1,
			Register: 2003,
			Value:    "1025",
			Type:     "int64",
			Alias:    "hello",
		},
	}
)

func main() {
	pub()
}

func pub() {

	// pack
	buf := &bytes.Buffer{}
	enc := codec.NewEncoder(buf, mh)
	enc.Encode(command)

	// zmq
	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Connect("ipc:///tmp/dummy")

	for {
		sender.Send("mbtcp.once.write", zmq.SNDMORE)
		sender.Send(buf.String(), 0) // convert buffer arr to str
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
