package main

import (
	_ "encoding/json"
	"fmt"
	proto "github.com/golang/protobuf/proto"
	zmq "github.com/taka-wang/zmq3"
	"time"
)

func main() {
	pub()
}

func pub() {
	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Connect("ipc:///tmp/dummy")

	command := &MbTcpMultipleWriteReq{ // named key
		CmdHeader: &CmdHeader{
			Receiver: "core",
			Sender:   "me",
			Version:  1,
			Tid:      32,
			Method:   "mbtcp.once.write",
		},
		MbTcpHeader: &MbTcpHeader{
			Ip:   "192.168.1.1",
			Port: 503,
			Id:   22,
		},
	}

	command.Requests = append(command.Requests,
		&MbWriteRequest{
			Code:     1,
			Register: 2003,
			Value:    "1025",
			Type:     "int64",
			Alias:    "hello",
		},
	)
	command.Requests = append(command.Requests,
		&MbWriteRequest{
			Code:     2,
			Register: 3003,
			Value:    "2025",
			Type:     "int32",
			Alias:    "hello_2",
		},
	)

	cmd, err := proto.Marshal(command)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}

	for {
		sender.Send("mbtcp.once.write", zmq.SNDMORE)
		sender.Send(string(cmd), 0)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
