package main

import (
	"encoding/json"
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

	command := &MbTcpHeader{
		Ip:   "192.168.1.1",
		Port: 503,
		Id:   22,
	}

	//test
	cmd1, _ := json.Marshal(command)
	fmt.Println(string(cmd1))

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
