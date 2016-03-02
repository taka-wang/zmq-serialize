package main

type MbTcpHeader struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Id   int    `json:"id"`
}
