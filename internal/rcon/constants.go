/*
Package rcon has the protocol specific implementations
*/
package rcon

const (
    // ServerDataAuth represents an auth request from client
    ServerDataAuth ServerDataType = 3
    // ServerDataAuthResponse represents an auth response from server
    ServerDataAuthResponse ServerDataType = 2
    // ServerDataExecCommand represents a command request from client
    ServerDataExecCommand ServerDataType = 2
    //ServerDataResponseValue represents a command response from server
    ServerDataResponseValue ServerDataType = 0

    // EmptyByte is a simple empty byte to be added at the end of packet message
    EmptyByte = 0x00

    // PacketPaddingSize is the packet padding
    PacketPaddingSize int32 = 2
    // PacketHeaderSize is the packet header size
    PacketHeaderSize int32 = 8
)
