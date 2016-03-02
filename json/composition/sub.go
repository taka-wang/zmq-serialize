package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/taka-wang/zmq3"
	_ "time"
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
		var s MbTcpSingleWriteReq
		//_, err := s.UnmarshalMsg([]byte(msg[1]))
		err := json.Unmarshal([]byte(msg[1]), &s) // frame 2
		if err != nil {
			fmt.Println(err) // unmarshal from json string failed
		} else {
			fmt.Println(msg[0])
			fmt.Println(s)
		}
	}
}
