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
        zframe_t *first = zmsg_pop(msg);
        zframe_t *payload = zmsg_pop(msg);
        char *first_buffer = zframe_strdup(first);
        char *payload_buffer = zframe_strdup(payload);
        printf("method: %s\n", first_buffer);
        //printf("payload: %s\n", payload_buffer);

        // unpack
        MbTcpMultipleWriteReq *command;
        
        CmdHeader *cmd_header;
        MbTcpHeader *mb_tcp_header;
        
        MbWriteRequest **mb_write_requests; // note!

        unsigned len = zframe_size(payload);

        command = mb_tcp_multiple_write_req__unpack(NULL, len, payload_buffer);
        cmd_header       = command->cmd_header;
        mb_tcp_header    = command->mb_tcp_header;
        mb_write_requests= command->requests; // note!
        
        if (command == NULL) {
            fprintf(stderr, "error unpacking incoming message\n");
            break;
        }

        printf("Recv: %s %s %d %lld %s\n", cmd_header->receiver, cmd_header->sender, cmd_header->version, cmd_header->tid, cmd_header->method);
        printf("Recv: ip: %s, port: %d, id: %d\n", mb_tcp_header->ip, mb_tcp_header->port, mb_tcp_header->id);
        
        int i = 0;
        for (i = 0; i < command->n_requests; ++i)
        {
            /*
            printf("Recv: %d %d %s %s %s\n",
                command->requests[i]->code, 
                command->requests[i]->register_, 
                command->requests[i]->value, 
                command->requests[i]->type, 
                command->requests[i]->alias);
            */
            printf("Recv: %d %d %s %s %s\n",
                mb_write_requests[i]->code, 
                mb_write_requests[i]->register_, 
                mb_write_requests[i]->value, 
                mb_write_requests[i]->type, 
                mb_write_requests[i]->alias);
        }

        //zmsg_dump(msg);
        
        //cleanup
        mb_tcp_multiple_write_req__free_unpacked(command, NULL); // free unpacked command
        zframe_destroy(&first);
        zframe_destroy(&payload);
        zmsg_destroy(&msg );
    }
    
    // cleanup
    zctx_destroy(&context);
    return 0;
}