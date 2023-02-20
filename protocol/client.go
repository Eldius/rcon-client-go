package protocol

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type ClientOptions struct {
	host   string
	currID int32
	conn   net.Conn
	debug  bool
}

type Options func(*ClientOptions)

/*
WithHost sets client host
*/
func WithHost(host string) Options {
	return func(opts *ClientOptions) {
		opts.host = strings.Trim(host, " ")
	}
}

/*
WithDebugLog sets client debug mode
*/
func WithDebugLog(debug bool) Options {
	return func(opts *ClientOptions) {
		opts.debug = debug
	}
}

/*
WithID sets client ID
*/
func WithID(id int32) Options {
	return func(opts *ClientOptions) {
		opts.currID = id
	}
}

// Client is the client implementation
type Client struct {
	opt  *ClientOptions
	conn net.Conn
}

// NewClient creates a new Client
func NewClient(cfg ...Options) (*Client, error) {
	o := ClientOptions{}
	for _, c := range cfg {
		c(&o)
	}
	if o.host == "" {
		return nil, errors.New("host must not be empty")
	}
	conn, err := net.Dial("tcp", o.host)
	if err != nil {
		return nil, err
	}
	return &Client{
		opt:  &o,
		conn: conn,
	}, nil
}

// Close disconnects from server
func (c *Client) Close() error {
	return c.conn.Close()
}

// Command executes a command
func (c *Client) Command(cmd string) (*Packet, error) {
	p := newPecket(c.nextID(), RequestTypeCommandRequest, cmd)
	return c.exec(p)
}

// Login logs in to the server
func (c *Client) Login(passwd string) (*Packet, error) {
	p := newPecket(c.nextID(), RequestTypeAuthRequest, passwd)
	return c.exec(p)
}

/*
readPacket reads data from server

	## Response packet structure
	Size        32-bit little-endian Signed Integer	Varies, see below.
	ID          32-bit little-endian Signed Integer	Varies, see below.
	Type        32-bit little-endian Signed Integer	Varies, see below.
	Body        Null-terminated ASCII String	Varies, see below.
	Empty       String	Null-terminated ASCII String	0x00
*/
func (c *Client) readPacket(p *Packet) error {
	if err := binary.Read(c.conn, binary.LittleEndian, &p.ResponseSize); err != nil {
		return fmt.Errorf("protocol: read Packet size: %w", err)
	}
	if err := binary.Read(c.conn, binary.LittleEndian, &p.ResponseID); err != nil {
		return fmt.Errorf("protocol: read Packet size: %w", err)
	}
	if err := binary.Read(c.conn, binary.LittleEndian, &p.ResponseType); err != nil {
		return fmt.Errorf("protocol: read Packet size: %w", err)
	}

	expectedBodySize := p.ResponseSize - PacketHeaderSize
	p.ResponseBody = make([]byte, expectedBodySize)

	var actualBodySize int32
	for actualBodySize < expectedBodySize {
		var m int
		m, err := c.conn.Read(p.ResponseBody[actualBodySize:])
		if err != nil {
			return fmt.Errorf("protocol: %w", err)
		}
		actualBodySize += int32(m)
	}
	c.debugLog(
		fmt.Sprintf("- response size: %d", p.ResponseSize),
		fmt.Sprintf("- response id: %d", p.ResponseID),
		fmt.Sprintf("- response type: %d", p.ResponseType),
		fmt.Sprintf("- expected body size: %d", expectedBodySize))
	//fmt.Sprintf("- response body: %s", p.ResponseBody),
	//fmt.Sprintf("- response body: %v", p.ResponseBody))

	return nil
}

func (c *Client) exec(p *Packet) (*Packet, error) {
	if err := c.writePacket(p); err != nil {
		return nil, err
	}
	if err := c.readPacket(p); err != nil {
		return nil, err
	}
	if !p.IsValid() {
		return p, fmt.Errorf("invalid response: [id: %d => %d] [type: '%d' => '%d']", p.ID, p.ResponseID, p.Type, p.ResponseType)
	}
	return p, nil
}

func (c *Client) writePacket(p *Packet) error {
	enc, err := p.ToBytes()
	if err != nil {
		return err
	}

	c.debugLog(
		fmt.Sprintf("[string] sending msg: '%s'", enc),
		fmt.Sprintf("[byte]   sending msg: %v", enc))

	_, err = c.conn.Write(enc)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) nextID() int32 {
	id := c.opt.currID
	c.opt.currID++
	c.debugLog(fmt.Sprintf("current id: %d => next id: %d", id, c.opt.currID))
	return id
}

func (c *Client) debugLog(msgs ...string) {
	if c.opt.debug {
		log.Println("[DEBUG] -- debug -----")
		for _, msg := range msgs {
			log.Printf("[DEBUG] %s\n", msg)
		}
		log.Println("[DEBUG] --------------")
	}
}
