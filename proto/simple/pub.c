// compile: gcc type.pb-c.h type.pb-c.c pub.c -lzmq -lczmq -lprotobuf-c  -o pubc

#include <stdio.h>
#include <stdlib.h>
#include "type.pb-c.h"
#include <czmq.h>
#include <unistd.h>

int main (int argc, char *argv [])
{

    Main__MbTcpHeader command = MAIN__MB_TCP_HEADER__INIT; // construct
    void *buf;                     // Buffer to store serialized data
    unsigned len;

    
    command.port = 503;
    command.id = 22;
    command.ip = "192.168.1.1";
    len = main__mb_tcp_header__get_packed_size(&command);
    buf = malloc(len);
    main__mb_tcp_header__pack(&command, buf);
    //fprintf(stderr, "Writing %d serialized bytes\n", len);
    //fwrite(buf,len,1,stdout);

    zctx_t *context = zctx_new();
    void *publisher = zsocket_new(context, ZMQ_PUB);
    zsocket_connect (publisher, "ipc:///tmp/dummy");
    while (true) {
        printf(".\n");
        zmsg_t *msg = zmsg_new();
        zmsg_addstr(msg, "mbtcp.once.write");         // frame 1
        //zmsg_addstr(msg, "{\"receiver\": \"core\"}"); // frame 2
        zmsg_addstr(msg, (char*)buf); // frame 2
        zmsg_send(&msg, publisher);
        //zclock_sleep(3 * 1000);
        sleep(3);
        zmsg_destroy(&msg);
    }
    
    // cleanup
    free(buf);
    zctx_destroy(&context);
    return 0;
}