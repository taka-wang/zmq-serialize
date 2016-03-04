package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/taka-wang/zmq3"
	"time"
)

var (
	command = MbTcpMultipleWriteReq{ // named key
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
		Requests: []MbWriteRequest{
			MbWriteRequest{
				Code:     1,
				Register: 2003,
				Value:    "1025",
				Type:     "int64",
				Alias:    "hello",
			},
			MbWriteRequest{
				Code:     1,
				Register: 2003,
				Value:    "1025",
				Type:     "int64",
				Alias:    "hello",
			},
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
		sender.Send(string(cmd), 0)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
