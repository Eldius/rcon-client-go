package rcon

const (
    ServerDataAuth          ServerDataType = 3
    ServerDataAuthResponse  ServerDataType = 2
    ServerDataExecCommand   ServerDataType = 2
    ServerDataResponseValue ServerDataType = 0

    EmptyByte = byte(0)
)

const (
    PacketPaddingSize int32 = 2 // Size of Packet's padding.
    PacketHeaderSize  int32 = 8 // Size of Packet's header.

    MinPacketSize = PacketPaddingSize + PacketHeaderSize
    MaxPacketSize = 4096 + MinPacketSize
)
