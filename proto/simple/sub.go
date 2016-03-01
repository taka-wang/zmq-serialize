package main

import (
	_ "encoding/json"
	"fmt"
	proto "github.com/golang/protobuf/proto"
	zmq "github.com/taka-wang/zmq3"
	_ "time"
)

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

		// decode
		s := &MbTcpHeader{}
		err := proto.Unmarshal([]byte(msg[1]), s)

		if err != nil {
			fmt.Println(err) // unmarshal from json string failed
		} else {
			fmt.Println(msg[0])
			fmt.Println(s)
		}
	}
}
