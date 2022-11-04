/*
Package rcon has the protocol specific implementations
*/
package rcon

const (
	// RequestTypeAuthRequest represents an auth request from client
	RequestTypeAuthRequest RequestType = 3
	// ResponseTypeAuth represents an auth response from server
	ResponseTypeAuth ResponseType = 2
	// RequestTypeCommandRequest represents a command request from client
	RequestTypeCommandRequest RequestType = 2
	//ResponseTypeCommand represents a command response from server
	ResponseTypeCommand ResponseType = 0

	// EmptyByte is a simple empty byte to be added at the end of packet message
	EmptyByte = 0x00

	// PacketPaddingSize is the packet padding
	PacketPaddingSize int32 = 2
	// PacketHeaderSize is the packet header size
	PacketHeaderSize int32 = 8
)
