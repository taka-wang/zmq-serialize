/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: type.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "type.pb-c.h"
void   main__mb_tcp_header__init
                     (Main__MbTcpHeader         *message)
{
  static Main__MbTcpHeader init_value = MAIN__MB_TCP_HEADER__INIT;
  *message = init_value;
}
size_t main__mb_tcp_header__get_packed_size
                     (const Main__MbTcpHeader *message)
{
  assert(message->base.descriptor == &main__mb_tcp_header__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t main__mb_tcp_header__pack
                     (const Main__MbTcpHeader *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &main__mb_tcp_header__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t main__mb_tcp_header__pack_to_buffer
                     (const Main__MbTcpHeader *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &main__mb_tcp_header__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Main__MbTcpHeader *
       main__mb_tcp_header__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Main__MbTcpHeader *)
     protobuf_c_message_unpack (&main__mb_tcp_header__descriptor,
                                allocator, len, data);
}
void   main__mb_tcp_header__free_unpacked
                     (Main__MbTcpHeader *message,
                      ProtobufCAllocator *allocator)
{
  assert(message->base.descriptor == &main__mb_tcp_header__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor main__mb_tcp_header__field_descriptors[3] =
{
  {
    "ip",
    1,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Main__MbTcpHeader, ip),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "port",
    2,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT32,
    offsetof(Main__MbTcpHeader, has_port),
    offsetof(Main__MbTcpHeader, port),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "id",
    3,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT32,
    offsetof(Main__MbTcpHeader, has_id),
    offsetof(Main__MbTcpHeader, id),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned main__mb_tcp_header__field_indices_by_name[] = {
  2,   /* field[2] = id */
  0,   /* field[0] = ip */
  1,   /* field[1] = port */
};
static const ProtobufCIntRange main__mb_tcp_header__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 3 }
};
const ProtobufCMessageDescriptor main__mb_tcp_header__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "main.MbTcpHeader",
  "MbTcpHeader",
  "Main__MbTcpHeader",
  "main",
  sizeof(Main__MbTcpHeader),
  3,
  main__mb_tcp_header__field_descriptors,
  main__mb_tcp_header__field_indices_by_name,
  1,  main__mb_tcp_header__number_ranges,
  (ProtobufCMessageInit) main__mb_tcp_header__init,
  NULL,NULL,NULL    /* reserved[123] */
};
