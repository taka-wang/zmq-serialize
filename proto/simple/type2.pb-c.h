/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: type2.proto */

#ifndef PROTOBUF_C_type2_2eproto__INCLUDED
#define PROTOBUF_C_type2_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1000000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1002001 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _Main__MbTcpHeader Main__MbTcpHeader;


/* --- enums --- */


/* --- messages --- */

struct  _Main__MbTcpHeader
{
  ProtobufCMessage base;
  char *ip;
  uint32_t port;
  uint32_t id;
};
#define MAIN__MB_TCP_HEADER__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&main__mb_tcp_header__descriptor) \
    , NULL, 0, 0 }


/* Main__MbTcpHeader methods */
void   main__mb_tcp_header__init
                     (Main__MbTcpHeader         *message);
size_t main__mb_tcp_header__get_packed_size
                     (const Main__MbTcpHeader   *message);
size_t main__mb_tcp_header__pack
                     (const Main__MbTcpHeader   *message,
                      uint8_t             *out);
size_t main__mb_tcp_header__pack_to_buffer
                     (const Main__MbTcpHeader   *message,
                      ProtobufCBuffer     *buffer);
Main__MbTcpHeader *
       main__mb_tcp_header__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   main__mb_tcp_header__free_unpacked
                     (Main__MbTcpHeader *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*Main__MbTcpHeader_Closure)
                 (const Main__MbTcpHeader *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor main__mb_tcp_header__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_type2_2eproto__INCLUDED */