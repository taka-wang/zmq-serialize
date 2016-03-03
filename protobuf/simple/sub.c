// compile: gcc type2.pb-c.h type2.pb-c.c sub.c -lzmq -lczmq -lprotobuf-c  -o subc

#include <czmq.h>
#include "type2.pb-c.h"

int main (int argc, char *argv [])
{
    zctx_t *context = zctx_new ();
    void *subscriber = zsocket_new (context, ZMQ_SUB);
    zsocket_bind (subscriber, "ipc:///tmp/dummy");

    zsocket_set_subscribe (subscriber, "mbtcp"); // set filter
    
    while (!zctx_interrupted) {
        zmsg_t *msg = zmsg_recv(subscriber);
        //zmsg_dump(msg); // debug
        zframe_t *first = zmsg_pop(msg);
        zframe_t *payload = zmsg_pop(msg);
        char *first_buffer = zframe_strdup(first);
        char *payload_buffer = zframe_strdup(payload);
        printf("method: %s\n", first_buffer);
        //printf("payload: %s\n", payload_buffer);

        // unpack
        Main__MbTcpHeader *command;
        unsigned len = zframe_size(payload);

        command = main__mb_tcp_header__unpack(NULL, len, payload_buffer);
        
        if (command == NULL) {
            fprintf(stderr, "error unpacking incoming message\n");
            break;
        }

        printf("Recv: ip: %s, port: %d, id: %d\n", command->ip, command->port, command->id);
        
        //cleanup
        main__mb_tcp_header__free_unpacked(command, NULL); // free unpacked command
        zframe_destroy(&first);
        zframe_destroy(&payload);
        zmsg_destroy(&msg );
    }
    
    // cleanup
    zctx_destroy(&context);
    return 0;
}