// compile: gcc type2.pb-c.h type2.pb-c.c pub.c -lzmq -lczmq -lprotobuf-c  -o pubc

#include <stdio.h>
#include <stdlib.h>
#include "type2.pb-c.h"
#include <czmq.h>
#include <unistd.h>


int main (int argc, char *argv [])
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
    cmd_header.version  = "mbtcp";
    cmd_header.tid      = "mbtcp";
    cmd_header.method   = "mbtcp";
    //TODO!!!!

    //command.has_port = 1;
    //command.has_id = 1;
    command.ip = "192.168.1.1";
    command.port = 503;
    command.id = 22;
    
    len = main__mb_tcp_header__get_packed_size(&command);
    
    buf = malloc(len);
    main__mb_tcp_header__pack(&command, buf);

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