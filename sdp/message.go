package sdp

import (
	"bytes"
	"strconv"
)

// ErrorSDPParsing returned when failed to parse SDP message
var ErrorSDPParsing = errorNew("Error parding SDP message")

// Message SDP message structure
type Message struct {
	ver     byte
	subject []byte
	Origin  Origin
	Conn    Conn
	Time    []TimeDesc
	Attr    []Attribute
	Medias  Medias
}

// Version SDP message version field
func (m *Message) Version() int {
	return int(m.ver - 0x30)
}

// Subject SDP message subject field
func (m *Message) Subject() string {
	return string(bytes.TrimSpace(m.subject))
}

// Origin SDP origin field (RFC4566 #5.2)
type Origin struct {
	username []byte
	sessID   []byte
	sessVer  []byte
	netType  []byte
	addrType []byte
	unicAddr []byte
}

// Username SDP origin field username
func (o Origin) Username() string {
	return string(o.username)
}

// SessionID SDP origin field session id
func (o Origin) SessionID() int {
	id, err := strconv.Atoi(string(o.sessID))
	if err != nil {
		return -1
	}
	return id
}

// SessionVer SDP origin field session version
func (o Origin) SessionVer() int {
	ver, err := strconv.Atoi(string(o.sessVer))
	if err != nil {
		return -1
	}
	return ver
}

// NetType SDP origin field net type
func (o Origin) NetType() string {
	return string(o.netType)
}

// AddrType SDP origin field address type
func (o Origin) AddrType() string {
	return string(o.addrType)
}

// UnicastAddr SDP origin field unicast address
func (o Origin) UnicastAddr() string {
	return string(o.unicAddr)
}

// Conn SDP connection data field (RFC4566 #5.7)
type Conn struct {
	netType  []byte
	addrType []byte
	address  []byte
}

// NetType SDP connection data field net type
func (c Conn) NetType() string {
	return string(c.netType)
}

// AddrType SDP connection data field address type
func (c Conn) AddrType() string {
	return string(c.addrType)
}

// Address SDP connection data field unicast address
func (c Conn) Address() string {
	return string(c.address)
}

// TimeDesc time description structure that contains time and repeat time fields
type TimeDesc struct {
	start []byte
	stop  []byte
}

// StartTime time description field
func (t TimeDesc) StartTime() int {
	time, err := strconv.Atoi(string(t.start))
	if err != nil {
		return -1
	}
	return time
}

// StopTime time description field
func (t TimeDesc) StopTime() int {
	time, err := strconv.Atoi(string(t.stop))
	if err != nil {
		return -1
	}
	return time
}

// Medias list of session medias
type Medias []Media

// Media media description of SDP session
type Media struct {
	mtype []byte
	port  []byte
	nport []byte
	proto []byte
	fmt   []byte
	attr  []Attribute
}

// Type SDP media field type
func (m Media) Type() string {
	return string(m.mtype)
}

// Port SDP media field port
func (m Media) Port() int {
	port, err := strconv.Atoi(string(m.port))
	if err != nil {
		return -1
	}
	return port
}

// NumPort SDP media field ports number
func (m Media) NumPort() int {
	n, err := strconv.Atoi(string(m.nport))
	if err != nil {
		return 0
	}
	return n
}

// Proto SDP media field protocol
func (m Media) Proto() string {
	return string(m.proto)
}

// Fmt SDP media field formats list
func (m Media) Fmt() string {
	return string(bytes.TrimSpace(m.fmt))
}

// Attribute SDP field
type Attribute struct {
	isFlag bool
	key    []byte
	value  []byte
	flag   []byte
}

// IsFlag returns true if attribute is flag (a=sendonly)
func (a Attribute) IsFlag() bool { return a.isFlag }

// Key attribute key (a=key:value)
func (a Attribute) Key() string { return string(a.key) }

// Value attribute value (a=key:value)
func (a Attribute) Value() string { return string(a.value) }

// Flag attribute (a=flag)
func (a Attribute) Flag() string { return string(a.flag) }
