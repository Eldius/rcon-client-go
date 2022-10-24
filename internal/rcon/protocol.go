package rcon

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "log"
    "net"
)

// ServerDataType represents the request type
type ServerDataType int32

// Client is the client implementation
type Client struct {
    host   string
    currID int32
    conn   net.Conn
}

// Packet is the command packet (including request and response data
type Packet struct {
    Size         int32
    Type         ServerDataType
    ID           int32
    Body         []byte
    ResponseSize int32
    ResponseType ServerDataType
    ResponseID   int32
    ResponseBody []byte
}

func newPecket(id int32, t ServerDataType, b string) *Packet {
    body := []byte(b)
    size := len(body) + int(PacketHeaderSize+PacketPaddingSize)
    return &Packet{
        Size: int32(size),
        Type: t,
        ID:   id,
        Body: body,
    }
}

// NewClient creates a new Client
func NewClient(host string) (*Client, error) {
    conn, err := net.Dial("tcp", host)
    if err != nil {
        return nil, err
    }
    return &Client{
        host:   host,
        currID: 1,
        conn:   conn,
    }, nil
}

// Command executes a command
func (c *Client) Command(cmd string) (*Packet, error) {
    p := newPecket(c.nextID(), ServerDataExecCommand, cmd)
    return c.exec(p)
}

// Login logs in to the server
func (c *Client) Login(passwd string) (*Packet, error) {
    p := newPecket(c.nextID(), ServerDataAuth, passwd)
    return c.exec(p)
}

func (c *Client) exec(p *Packet) (*Packet, error) {
    if err := c.sendPacket(p); err != nil {
        return nil, err
    }
    if err := c.Read(p); err != nil {
        return nil, err
    }
    return p, nil
}

func (c *Client) sendPacket(p *Packet) error {
    enc, err := p.ToBytes()
    if err != nil {
        return err
    }
    log.Printf("[string] sending msg: '%s'\n", enc)
    log.Printf("[byte]   sending msg: %v\n", enc)
    _, err = c.conn.Write(enc)
    if err != nil {
        return err
    }
    return nil
}

func (c *Client) nextID() int32 {
    id := c.currID
    c.currID++
    return id
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
Read reads data from server
  ## Response packet structure
  Size        32-bit little-endian Signed Integer	Varies, see below.
  ID          32-bit little-endian Signed Integer	Varies, see below.
  Type        32-bit little-endian Signed Integer	Varies, see below.
  Body        Null-terminated ASCII String	Varies, see below.
  Empty       String	Null-terminated ASCII String	0x00
*/
func (c *Client) Read(p *Packet) error {
    if err := binary.Read(c.conn, binary.LittleEndian, &p.ResponseSize); err != nil {
        return fmt.Errorf("rcon: read Packet size: %w", err)
    }
    log.Printf("- response size: %d\n", p.ResponseSize)

    if err := binary.Read(c.conn, binary.LittleEndian, &p.ResponseID); err != nil {
        return fmt.Errorf("rcon: read Packet size: %w", err)
    }
    log.Printf("- response id: %d\n", p.ResponseID)

    if err := binary.Read(c.conn, binary.LittleEndian, &p.ResponseType); err != nil {
        return fmt.Errorf("rcon: read Packet size: %w", err)
    }
    log.Printf("- response type: %d\n", p.ResponseType)

    expectedBodySize := p.ResponseSize - PacketHeaderSize
    p.ResponseBody = make([]byte, expectedBodySize)

    log.Printf("- expected body size: %d\n", expectedBodySize)
    var actualBodySize int32
    for actualBodySize < expectedBodySize {
        var m int
        m, err := c.conn.Read(p.ResponseBody[actualBodySize:])
        if err != nil {
            return fmt.Errorf("rcon: %w", err)
        }
        actualBodySize += int32(m)
    }
    log.Printf("- response body: %s\n", p.ResponseBody)
    log.Printf("- response body: %v\n", p.ResponseBody)

    return nil
}
