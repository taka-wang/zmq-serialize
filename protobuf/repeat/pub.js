var zmq = require('zmq')
, protobuf = require('protocol-buffers')
, fs = require('fs')
, messages = protobuf(fs.readFileSync('type.proto'))
, pub = zmq.socket('pub')
, command = {
    "cmd_header":{
        "receiver": "mbtcp",
        "sender":   "restful",        
        "version":  1,
        "tid":      33,                
        "method":   "mbtcp.once.write"
    },
    "mb_tcp_header":{
        "ip": "192.168.1.1", 
        "port": 503,
        "id": 22
    },
    "requests": [ 
        {
            "code":     1,                 
            "register": 2003,             
            "value":    "1025",            
            "type":     "int64",        
            "alias":    "hello_1"  
        }, 
        {
            "code":     2,                 
            "register": 3003,             
            "value":    "2025",            
            "type":     "int32",        
            "alias":    "hello_2"  
        }
    ]
}
, buf = messages.MbTcpMultipleWriteReq.encode(command); // encode

pub.connect("ipc:///tmp/dummy"); // connect to zmq endpoint

setInterval(function() {
    pub.send("mbtcp.once.write", zmq.ZMQ_SNDMORE);
    pub.send(buf);
}, 1000); // emit every 1 seconds
