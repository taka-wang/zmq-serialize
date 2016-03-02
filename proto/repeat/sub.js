var zmq = require('zmq')
, protobuf = require('protocol-buffers')
, fs = require('fs')
, messages = protobuf(fs.readFileSync('type.proto'))
, sub = zmq.socket('sub')

sub.bindSync("ipc:///tmp/dummy"); 	// bind to zmq endpoint
sub.subscribe("mbtcp"); 			// filter topic

sub.on("message", function(method, payload) {
    console.log(method.toString());
    var obj = messages.MbTcpMultipleWriteReq.decode(payload); //decode
    console.dir(obj);
});