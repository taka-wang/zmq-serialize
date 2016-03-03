// compile: gcc type.pb-c.h type.pb-c.c sub.c -lzmq -lczmq -lprotobuf-c  -o subc

#include <czmq.h>
#include "type.pb-c.h"

int main (int argc, char *argv [])
{
    zctx_t *context = zctx_new ();
    void *subscriber = zsocket_new (context, ZMQ_SUB);
    zsocket_bind (subscriber, "ipc:///tmp/dummy");

    zsocket_set_subscribe (subscriber, "mbtcp"); // set filter
    
    while (!zctx_interrupted) {
        zmsg_t *msg = zmsg_recv(subscriber);
        zframe_t *first = zmsg_pop(msg);
        zframe_t *payload = zmsg_pop(msg);
        char *first_buffer = zframe_strdup(first);
        char *payload_buffer = zframe_strdup(payload);
        printf("method: %s\n", first_buffer);
        //printf("payload: %s\n", payload_buffer);

        uint8_t buffer = zframe_get_uint8(payload);
        printf("%u\n", buffer);

        //--
        //Main__MbTcpHeader *command;
        //command = amessage__unpack(NULL, msg_len, buf);





        //zmsg_dump(msg);
        zframe_destroy(&first);
        zframe_destroy(&payload);
        zmsg_destroy(&msg );
    }
    
    // cleanup
    zctx_destroy(&context);
    return 0;
}