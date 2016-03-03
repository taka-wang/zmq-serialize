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
    
    MbWriteRequest **mb_write_requests; // without init
    
    MbTcpMultipleWriteReq command = MB_TCP_MULTIPLE_WRITE_REQ__INIT;

    cmd_header.receiver = "mbtcp";
    cmd_header.sender   = "restful";
    cmd_header.version  = 1;
    cmd_header.tid      = 33;
    cmd_header.method   = "mbtcp.once.write";
    
    mb_tcp_header.ip    = "192.168.1.1";
    mb_tcp_header.port  = 503;
    mb_tcp_header.id    = 22;
    
    command.n_requests  = 2;
    mb_write_requests = malloc (sizeof (MbWriteRequest*) * command.n_requests);

    int i = 0;
    for (i = 0; i < command.n_requests; ++i)
    {
        mb_write_requests[i] = malloc (sizeof (MbWriteRequest));
        mb_write_request__init (mb_write_requests[i]); // note!! dynamic init

        // assignment
        mb_write_requests[i]->code       = 1;
        mb_write_requests[i]->register_  = 2003;        // strange code gen for "register"
        mb_write_requests[i]->value      = "1025";
        mb_write_requests[i]->type       = "int64";
        mb_write_requests[i]->alias      = "hello";
    }

    command.cmd_header      = &cmd_header;
    command.mb_tcp_header   = &mb_tcp_header;
    command.requests        = mb_write_requests;    // note!

    len = mb_tcp_multiple_write_req__get_packed_size (&command);
    buf = malloc (len);                     // Allocate memory
    mb_tcp_multiple_write_req__pack (&command, buf);

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
    // free repeat
    for (i = 0; i < command.n_requests; ++i) {
        free(mb_write_requests[i]);
    }
    free(mb_write_requests);

    zctx_destroy(&context);
    return 0;
}