var zmq = require('zmq')
, msgpack = require('msgpack5')()
, encode  = msgpack.encode
, decode  = msgpack.decode
, pub = zmq.socket('pub')
, command = {
    "cmd_header":{
        "receiver": "mbtcp",
        "sender":   "restful",        
        "version":  "1",
        "tid":      33,                
        "method":   "mbtcp.once.write"
    },
    "mb_tcp_header":{
        "ip": "192.168.1.1", 
        "port": 503,
        "id": 22
    },
    "mb_write_request":{
        "code":     1,                 
        "register": 2003,             
        "value":    "1025",            
        "type":     "int64",        
        "alias":    "hello_1"
    }
}

pub.connect("ipc:///tmp/dummy"); // connect to zmq endpoint

setInterval(function() {
    pub.send("mbtcp.once.write", zmq.ZMQ_SNDMORE);
    pub.send(encode(command));
}, 1000); // emit every 1 seconds
