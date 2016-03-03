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

    Main__MbTcpHeader command = MAIN__MB_TCP_HEADER__INIT; // construct

    // assignment
    command.ip = "192.168.1.1";
    command.port = 503;
    command.id = 22;
    
    len = main__mb_tcp_header__get_packed_size(&command); // get length
    buf = malloc(len);                                    // allocate memory
    main__mb_tcp_header__pack(&command, buf);             // do pack

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
        zmsg_addstr(msg, (char*)buf);           // convert to char pointer: frame 2
        zmsg_send(&msg, publisher);
        zclock_sleep(1000); // sleep 1 second
        zmsg_destroy(&msg);
    }
    
    // cleanup
    free(buf);
    zctx_destroy(&context);
    return 0;
}