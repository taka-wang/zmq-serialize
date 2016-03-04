package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/taka-wang/zmq3"
	"time"
)

var (
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

	// marshal to json string
	cmd, err := json.Marshal(command)
	if err != nil {
		fmt.Println("json err:", err)
	}

	// zmq
	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Connect("ipc:///tmp/dummy")

	for {
		sender.Send("mbtcp.once.write", zmq.SNDMORE)
		sender.Send(string(cmd), 0) // byte arr to string
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
