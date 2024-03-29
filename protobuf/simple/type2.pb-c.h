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


typedef struct _MbTcpHeader MbTcpHeader;


/* --- enums --- */


/* --- messages --- */

struct  _MbTcpHeader
{
  ProtobufCMessage base;
  char *ip;
  uint32_t port;
  uint32_t id;
};
#define MB_TCP_HEADER__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mb_tcp_header__descriptor) \
    , NULL, 0, 0 }


/* MbTcpHeader methods */
void   mb_tcp_header__init
                     (MbTcpHeader         *message);
size_t mb_tcp_header__get_packed_size
                     (const MbTcpHeader   *message);
size_t mb_tcp_header__pack
                     (const MbTcpHeader   *message,
                      uint8_t             *out);
size_t mb_tcp_header__pack_to_buffer
                     (const MbTcpHeader   *message,
                      ProtobufCBuffer     *buffer);
MbTcpHeader *
       mb_tcp_header__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mb_tcp_header__free_unpacked
                     (MbTcpHeader *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*MbTcpHeader_Closure)
                 (const MbTcpHeader *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor mb_tcp_header__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_type2_2eproto__INCLUDED */