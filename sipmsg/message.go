package sipmsg

import (
	"bytes"
	"strconv"
)

type ptr uint16

// Structure to replresent position in []byte buffer
// "p" points to start position and "l" points to the last.
// That helps to avoid additional memory allocations.
type pl struct {
	p ptr
	l ptr
}

// ErrorSIPMsgParse SIP message parsing error
var ErrorSIPMsgParse = errorNew("Invalid SIP Message")

// ErrorSIPMsgCreate SIP message creating error
var ErrorSIPMsgCreate = errorNew("Invalid SIP Message")

// Message SIP message structure
type Message struct {
	ReqLine    *RequestLine
	StatusLine *StatusLine
	From       *HeaderFromTo
	To         *HeaderFromTo
	CSeq       *CSeq
	Contacts   ContactsList
	Vias       ViaList
	Routes     RouteList
	RecRoutes  RouteList
	CallID     string
	ContentLen uint // Content-Length
	Expires    uint
	MaxFwd     uint
	Headers    HeadersList
}

// MsgParse parser SIP message to Message structure
func MsgParse(data []byte) (*Message, error) {
	msg := &Message{}

	idx := bytes.Index(data, []byte("\r\n"))
	if idx == -1 {
		return nil, ErrorSIPMsgParse
	}
	// parse first line
	idx += 2
	hid, err := parseHeader(msg, data[:idx])
	if err != nil {
		return nil, err
	}
	if !(hid == SIPHdrRequestLine || hid == SIPHdrStatusLine) {
		return nil, ErrorSIPMsgParse.msg("Missing Request/Status line")
	}

	start := idx
	for i := idx; i < len(data); {
		if bytes.HasPrefix(data[i:], []byte("\r\n")) {
			i += 2
			if i < len(data) && (data[i] == ' ' || data[i] == '\t') {
				continue
			}
			hid, err = parseHeader(msg, data[start:i])
			if err != nil {
				return nil, err
			}
			if hid == MsgEOF {
				break
			}
			start = i
			continue
		}
		i++
	}
	// must be CRLF in the end of the SIP Message
	if hid != MsgEOF {
		return nil, ErrorSIPMsgParse.msg("Message must be finished with CRLF (%d)", hid)
	}
	return msg, nil
}

// NewRequest init basic SIP Request
func NewRequest(met, ruri string, via *Via, to, from *HeaderFromTo, cseq, maxfwd int) (*Message, error) {
	msg := &Message{}
	var buf []byte
	plName, plVal := pl{}, pl{}

	if cseq < 0 || cseq > (1<<31) {
		return nil, ErrorSIPMsgCreate.msg("CSeq value %d", cseq)
	}

	if maxfwd < 0 || maxfwd > 255 {
		return nil, ErrorSIPMsgCreate.msg("Max-Forwards value %d", cseq)
	}

	msg.ReqLine = NewReqLine(met, ruri)

	msg.Vias = append(msg.Vias, via)
	msg.pushHeader(SIPHdrVia, via.buf.Bytes(), via.name, pl{via.name.l + 2, via.buf.plen()})

	msg.From = from
	msg.From.AddTag()
	msg.pushHeader(SIPHdrFrom, from.buf.Bytes(), from.name, pl{from.name.l + 2, from.buf.plen()})

	msg.To = to
	msg.pushHeader(SIPHdrTo, to.buf.Bytes(), to.name, pl{to.name.l + 2, to.buf.plen()})

	msg.CallID = hashString()
	buf, plName, plVal = headerValue("Call-ID", msg.CallID)
	msg.pushHeader(SIPHdrCallID, buf, plName, plVal)

	msg.CSeq = &CSeq{uint(cseq), met}
	buf, plName, plVal = headerValue("CSeq", strconv.Itoa(cseq), met)
	msg.pushHeader(SIPHdrCSeq, buf, plName, plVal)

	msg.MaxFwd = uint(maxfwd)
	buf, plName, plVal = headerValue("Max-Forwards", strconv.Itoa(maxfwd))
	msg.pushHeader(SIPHdrMaxForwards, buf, plName, plVal)

	return msg, nil
}

// IsRequest returns true is SIP Message is request
func (m *Message) IsRequest() bool { return m.ReqLine != nil }

// IsResponse returns true is SIP Message is response
func (m *Message) IsResponse() bool { return m.StatusLine != nil }

// Bytes SIP message as bytes
func (m *Message) Bytes() []byte {
	b := m.buffer()
	return b.Bytes()
}

// String SIP message as string
func (m *Message) String() string {
	b := m.buffer()
	return b.String()
}

func (m *Message) buffer() buffer {
	var buf buffer
	if m.IsRequest() {
		buf.Write(m.ReqLine.Bytes())
	} else {
		buf.Write(m.StatusLine.Bytes())
	}

	for _, h := range m.Headers {
		buf.Write(h.buf)
	}
	buf.crlf()
	return buf
}

// private methods
func (m *Message) setStatusLine(buf []byte, pos []pl) HdrType {
	sl := &StatusLine{
		ver:    pos[0],
		code:   pos[1],
		reason: pos[2],
	}
	sl.buf.init(buf)
	m.StatusLine = sl
	return SIPHdrStatusLine
}

func (m *Message) setRequestLine(buf []byte, pos []pl) HdrType {
	rl := &RequestLine{
		method: pos[0],
		uri:    pos[1],
		ver:    pos[2],
	}
	rl.buf.init(buf)
	m.ReqLine = rl
	return SIPHdrRequestLine
}

func (m *Message) setCSeq(buf []byte, pos []pl) HdrType {
	num := buf[pos[1].p:pos[1].l]
	// do not check return. Parser must assure it is a number
	cseq, _ := strconv.ParseUint(string(num), 10, 32)
	m.CSeq = &CSeq{uint(cseq), string(buf[pos[2].p:pos[2].l])}

	m.pushHeader(SIPHdrCSeq, buf, pos[0], pl{pos[1].p, pos[2].l})
	return SIPHdrCSeq
}

func (m *Message) setCallID(buf []byte, pos []pl) HdrType {
	m.CallID = string(buf[pos[1].p:pos[1].l])
	m.pushHeader(SIPHdrCallID, buf, pos[0], pos[1])
	return SIPHdrCallID
}

func (m *Message) setContentLen(buf []byte, pos []pl) HdrType {
	num := buf[pos[1].p:pos[1].l]
	// do not check return. Parser must assure it is a number
	ln, _ := strconv.ParseUint(string(num), 10, 32)
	m.ContentLen = uint(ln)
	m.pushHeader(SIPHdrContentLength, buf, pos[0], pos[1])
	return SIPHdrContentLength
}

func (m *Message) setFrom(buf []byte, params []pl, fname, dname, addr, tag pl) HdrType {
	m.From = initHeaderFromTo(buf, params, fname, dname, addr, tag)
	if h := m.Headers.Find(SIPHdrFrom); h == nil {
		m.pushHeader(SIPHdrFrom, buf, fname, pl{fname.l + 1, ptr(len(buf))})
	}
	return SIPHdrFrom
}

func (m *Message) setTo(buf []byte, params []pl, fname, dname, addr, tag pl) HdrType {
	m.To = initHeaderFromTo(buf, params, fname, dname, addr, tag)
	if h := m.Headers.Find(SIPHdrTo); h == nil {
		m.pushHeader(SIPHdrTo, buf, fname, pl{fname.l + 1, ptr(len(buf))})
	}
	return SIPHdrTo
}

func (m *Message) setContact(buf []byte, name, dname, addr pl, params []pl, i int) {
	var b buffer
	b.init(buf)
	if m.Contacts.Count() == 0 || m.Contacts.Count() == i {
		m.Contacts.cnt = append(m.Contacts.cnt, &Contact{buf: b, name: name})
	}
	m.Contacts.cnt[i].name = name
	m.Contacts.cnt[i].dname = dname
	m.Contacts.cnt[i].addr = addr
	m.Contacts.cnt[i].params = params

	if !m.Headers.exists(buf) {
		m.pushHeader(SIPHdrContact, buf, name, pl{name.l + 1, b.plen()})
	}
}

func (m *Message) setContactStar() {
	m.Contacts.star = true
	m.pushHeader(SIPHdrContact, []byte("Contact: *\r\n"), pl{0, 7}, pl{9, 10})
}

func (m *Message) setVia(data []byte, name, trans, addr, port, branch, ttl, maddr, recevd pl, i int) {
	if m.Vias.Count() == 0 || m.Vias.Count() == i {
		var buf buffer
		buf.init(data)
		m.Vias = append(m.Vias, &Via{buf: buf, name: name})

		if !m.Headers.exists(data) {
			m.pushHeader(SIPHdrVia, buf.Bytes(), name, pl{name.l + 1, buf.plen()})
		}
	}
	m.Vias[i].trans = trans
	m.Vias[i].host = addr
	m.Vias[i].port = port
	m.Vias[i].branch = branch
	m.Vias[i].ttl = ttl
	m.Vias[i].maddr = maddr
	m.Vias[i].recevd = recevd
}

func (m *Message) setRoute(hid HdrType, buf []byte, fname, dname, addr pl, params []pl) {
	b := buffer{}
	b.init(buf)
	r := &Route{
		buf:    b,
		fname:  fname,
		dname:  dname,
		addr:   addr,
		params: params,
	}
	if hid == SIPHdrRecordRoute {
		m.RecRoutes = append(m.RecRoutes, r)
		m.pushHeader(SIPHdrRecordRoute, buf, fname, pl{fname.l + 1, r.buf.plen()})
		return
	}
	m.Routes = append(m.Routes, r)
	m.pushHeader(SIPHdrRoute, buf, fname, pl{fname.l + 1, r.buf.plen()})
}

func (m *Message) setExpires(num []byte) HdrType {
	// do not check return. Parser must assure it is a number
	expires, _ := strconv.ParseUint(string(num), 10, 32)
	m.Expires = uint(expires)
	return SIPHdrExpires
}

func (m *Message) setMaxFwd(num []byte) HdrType {
	// do not check return. Parser must assure it is a number
	max, _ := strconv.ParseUint(string(num), 10, 32)
	m.MaxFwd = uint(max)
	return SIPHdrMaxForwards
}

func (m *Message) setGenericHeader(buf []byte, pos []pl, id HdrType) HdrType {
	l := len(pos) - 1
	// non-determinism workarround
	// TODO: improve multiline value parsing (?)
	if int(pos[l].l) < len(buf)-2 {
		return -1
	}
	m.pushHeader(id, buf, pos[0], pos[l])
	return id
}

func (m *Message) pushHeader(id HdrType, buf []byte, name, value pl) {
	h := &Header{
		buf:   buf,
		id:    id,
		name:  name,
		value: value,
	}
	m.Headers = append(m.Headers, h)
}
