var zmq = require('zmq')
, sub = zmq.socket('sub')

sub.bindSync("ipc:///tmp/dummy"); 	// bind to zmq endpoint
sub.subscribe("mbtcp"); 			// filter topic

sub.on("message", function(method, payload) {
    console.log(method.toString());
    console.log(payload.toString());
});