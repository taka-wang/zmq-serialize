package main

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
	CmdHeader      `json:"cmd_header"`
	MbTcpHeader    `json:"mb_tcp_header"`
	MbWriteRequest `json:"mb_write_request"`
}
