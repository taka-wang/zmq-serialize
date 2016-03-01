var zmq = require('zmq')
, pub = zmq.socket('pub')
, command = {
    "ip": "192.168.1.1", 
    "port": 503,
    "id": 0 
}

pub.connect("ipc:///tmp/dummy"); // connect to zmq endpoint

setInterval(function() {
    pub.send("mbtcp.once.write", zmq.ZMQ_SNDMORE);
    pub.send(encode(command));
}, 1000); // emit every 1 seconds
