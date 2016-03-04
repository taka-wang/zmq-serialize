package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/taka-wang/zmq3"
	"time"
)

var (
	command = MbTcpHeader{
		Ip:   "192.168.1.1",
		Port: 503,
		Id:   22,
	}
)

func main() {
	pub()
}

func pub() {
	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Connect("ipc:///tmp/dummy")

	// marshal to json string
	cmd, err := json.Marshal(command)
	if err != nil {
		fmt.Println("json err:", err)
	}

	for {
		sender.Send("mbtcp.once.write", zmq.SNDMORE)
		sender.Send(string(cmd), 0) // byte arr to string
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
