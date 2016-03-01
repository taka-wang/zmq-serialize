package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/taka-wang/zmq3"
	"time"
)

type MbTcpHeader struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
	ID   int    `json:"id"`
}

func main() {
	pub()
}

func pub() {
	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Connect("ipc:///tmp/dummy")

	command := MbTcpHeader{
		IP:   "192.168.1.1",
		Port: 503,
		ID:   22,
	}

	// marshal to json string
	cmd, err := json.Marshal(command)
	if err != nil {
		fmt.Println("json err:", err)
	}

	for {
		sender.Send("mbtcp.once.write", zmq.SNDMORE)
		sender.Send(string(cmd), 0)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
