var zmq = require('zmq')
, protobuf = require('protocol-buffers')
, fs = require('fs')
, messages = protobuf(fs.readFileSync('type.proto'))
, pub = zmq.socket('pub')
, command = {
    "ip": "192.168.1.1", 
    "port": 503,
    "id": 22 
}
, buf = messages.MbTcpHeader.encode(command); // encode

pub.connect("ipc:///tmp/dummy"); // connect to zmq endpoint

setInterval(function() {
    pub.send("mbtcp.once.write", zmq.ZMQ_SNDMORE);
    pub.send(buf);
}, 1000); // emit every 1 seconds
