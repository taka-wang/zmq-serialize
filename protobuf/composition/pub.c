// compile: gcc type2.pb-c.h type2.pb-c.c pub.c -lzmq -lczmq -lprotobuf-c  -o pubc

#include <stdio.h>
#include <stdlib.h>
#include "type2.pb-c.h"
#include <czmq.h>
#include <unistd.h>

void* pack()
{
    void *buf; // Buffer to store serialized data
    unsigned len;

    // initialize
    CmdHeader cmd_header = CMD_HEADER__INIT;
    MbTcpHeader mb_tcp_header = MB_TCP_HEADER__INIT;
    MbWriteRequest mb_write_request = MB_WRITE_REQUEST__INIT;
    MbTcpSingleWriteReq command = MB_TCP_SINGLE_WRITE_REQ__INIT;

    cmd_header.receiver = "mbtcp";
    cmd_header.sender   = "restful";
    cmd_header.version  = 1;
    cmd_header.tid      = 33;
    cmd_header.method   = "mbtcp.once.write";
    
    mb_tcp_header.ip    = "192.168.1.1";
    mb_tcp_header.port  = 503;
    mb_tcp_header.id    = 22;
    
    mb_write_request.code       = 1;
    mb_write_request.register_  = 2003; // strange code gen
    mb_write_request.value      = "1025";
    mb_write_request.type       = "int64";
    mb_write_request.alias      = "hello";

    command.cmd_header  = &cmd_header;
    command.mb_tcp_header  = &mb_tcp_header;
    command.mb_write_request  = &mb_write_request;

    len = mb_tcp_single_write_req__get_packed_size (&command);
    buf = malloc (len);                     // Allocate memory
    mb_tcp_single_write_req__pack (&command, buf);

    return buf;
}

int main (int argc, char *argv [])
{

    void *buf = pack(); // pack message

    zctx_t *context = zctx_new();
    void *publisher = zsocket_new(context, ZMQ_PUB);
    zsocket_connect (publisher, "ipc:///tmp/dummy");
    
    while (!zctx_interrupted) { // handle ctrl+c
        zmsg_t *msg = zmsg_new();
        zmsg_addstr(msg, "mbtcp.once.write");   // frame 1
        zmsg_addstr(msg, (char*)buf);           // frame 2
        zmsg_send(&msg, publisher);
        zclock_sleep(1000); // sleep 1 second
        zmsg_destroy(&msg);
    }
    
    // cleanup
    free(buf);
    zctx_destroy(&context);
    return 0;
}