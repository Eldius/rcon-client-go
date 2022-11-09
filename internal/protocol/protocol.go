package protocol

import (
	"bytes"
	"embed"
	"encoding/binary"
	"fmt"
	"text/template"
)

//go:embed templates/*
var templateFiles embed.FS

var (
	tmpl *template.Template
)

func init() {
	tmpl = template.Must(template.ParseFS(templateFiles, "templates/*.tpl"))
}

// RequestType represents the client request type
type RequestType int32

// ResponseType represents the server response type
type ResponseType int32

// Packet is the command packet (including request and response data
type Packet struct {
	Size         int32
	Type         RequestType
	ID           int32
	Body         []byte
	ResponseSize int32
	ResponseType ResponseType
	ResponseID   int32
	ResponseBody []byte
}

func newPecket(id int32, t RequestType, b string) *Packet {
	body := []byte(b)
	size := len(body) + int(PacketHeaderSize+PacketPaddingSize)
	return &Packet{
		Size: int32(size),
		Type: t,
		ID:   id,
		Body: body,
	}
}

// TypeAsString returns a request type as a human-readable value
func (p *Packet) TypeAsString() string {
	switch p.Type {
	case RequestTypeAuthRequest:
		return "auth request"
	case RequestTypeCommandRequest:
		return "command request"
	default:
		return fmt.Sprintf("%d", p.Type)
	}
}

// ResponseTypeAsString returns a response type as a human-readable value
func (p *Packet) ResponseTypeAsString() string {
	switch p.ResponseType {
	case ResponseTypeAuth:
		return "auth response"
	case ResponseTypeCommand:
		return "command response"
	default:
		return fmt.Sprintf("%d", p.ResponseType)
	}
}

// BodyAsString returns request body content as string
func (p *Packet) BodyAsString() string {
	return string(p.Body)
}

// ResponseBodyAsString returns response body content as string
func (p *Packet) ResponseBodyAsString() string {
	return string(p.ResponseBody)
}

// String creates a string representation of this package
func (p *Packet) String() string {
	var buff bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buff, "packet.tpl", p); err != nil {
		panic(err)
	}
	return buff.String()
}

/*
ToBytes encodes the packet to send to server

	| Request packet structure |
	|--------|---------------------------------------------------------------|
	| Size   |      32-bit little-endian, Signed Integer	Varies, see below. |
	| ID     |      32-bit little-endian, Signed Integer	Varies, see below. |
	| Type   |      32-bit little-endian, Signed Integer	Varies, see below. |
	| Body   |      Null-terminated ASCII, String	Varies, see below.         |
	| Empty  |      String, Null-terminated, ASCII String	0x00               |
*/
func (p *Packet) ToBytes() ([]byte, error) {
	buffer := bytes.NewBuffer(make([]byte, 0, p.Size+4))

	_ = binary.Write(buffer, binary.LittleEndian, p.Size)
	_ = binary.Write(buffer, binary.LittleEndian, p.ID)
	_ = binary.Write(buffer, binary.LittleEndian, p.Type)

	// Write command body, null terminated ASCII string and an empty ASCIIZ string.
	buffer.Write(append(p.Body, EmptyByte, EmptyByte))
	return buffer.Bytes(), nil
}

/*
IsValid validates execution result
Validations:
- request and response have the same response ID
- response has a valid type for an auth request
- response has a valid type for a command execution request
*/
func (p *Packet) IsValid() bool {
	return (p.Type == RequestTypeAuthRequest && p.ResponseType == ResponseTypeAuth) || // has a valid auth response
		(p.Type == RequestTypeCommandRequest && p.ResponseType == ResponseTypeCommand) // has a valid execution response
}
