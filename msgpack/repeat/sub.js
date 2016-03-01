var zmq = require('zmq')
, msgpack = require('msgpack5')()
, encode  = msgpack.encode
, decode  = msgpack.decode
, sub = zmq.socket('sub')

sub.bindSync("ipc:///tmp/dummy"); 	// bind to zmq endpoint
sub.subscribe("mbtcp"); 			// filter topic

sub.on("message", function(method, payload) {
    console.log(method.toString());
    console.log(decode(payload));
});