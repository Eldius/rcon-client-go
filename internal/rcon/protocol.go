package rcon

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

type ServerDataType int32

type Client struct {
	host   string
	currID int32
}

type packet struct {
	Type ServerDataType
	ID   int32
	Body string
}

func NewClient(host string) *Client {
	return &Client{
		host:   host,
		currID: 1,
	}
}

func SendCommand(cmd string) error {
	return nil
}

func (c *Client) Login(passwd string) error {
	conn, err := net.Dial("tcp", c.host)
	if err != nil {
		return err
	}
	p := packet{
		Type: ServerDataAuth,
		ID:   c.nextID(),
		Body: passwd,
	}
	enc, err := p.ToBytes()
	if err != nil {
		return err
	}
	log.Printf("sending msg: '%s'\n", enc)
	n, err := conn.Write(enc)
	if err != nil {
		return err
	}
	log.Printf("wrote %d bytes to socket\n", n)
	var res []byte
	conn, err = net.Dial("tcp", c.host)
	if err != nil {
		return err
	}
	n, err = conn.Read(res)
	if err != nil {
		return err
	}
	log.Printf("response (%d): '%s'\n", n, res)
	return nil
}

func (c *Client) nextID() int32 {
	id := c.currID
	c.currID++
	return id
}

func (p *packet) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	size := len(p.Body)
	if err := binary.Write(buf, binary.LittleEndian, int32(size)); err != nil {
		log.Println("failed to encode size")
		return buf.Bytes(), err
	}
	if err := binary.Write(buf, binary.LittleEndian, p.ID); err != nil {
		log.Println("failed to encode id")
		return buf.Bytes(), err
	}
	if err := binary.Write(buf, binary.LittleEndian, p.Type); err != nil {
		log.Println("failed to encode type")
		return buf.Bytes(), err
	}
	if err := binary.Write(buf, binary.LittleEndian, []byte(p.Body)); err != nil {
		log.Println("failed to encode body")
		return buf.Bytes(), err
	}
	if err := binary.Write(buf, binary.LittleEndian, []byte(string(rune(0)))); err != nil {
		log.Println("failed to encode empty string")
		return buf.Bytes(), err
	}
	if err := binary.Write(buf, binary.LittleEndian, []byte(string(rune(0)))); err != nil {
		log.Println("failed to encode empty string")
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
