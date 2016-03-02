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

type MbTcpHeader struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Id   int    `json:"id"`
}

func main() {
	pub()
}

func pub() {
	sender, _ := zmq.NewSocket(zmq.PUB)
	defer sender.Close()
	sender.Connect("ipc:///tmp/dummy")

	command := MbTcpHeader{
		Ip:   "192.168.1.1",
		Port: 503,
		Id:   22,
	}

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
