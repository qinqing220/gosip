package sipmsg

import (
	"bytes"
	"strings"
)

// HdrType type header ID
type HdrType int

// SIP Header identifiers
const (
	MsgEOF HdrType = iota
	SIPHdrGeneric
	SIPHdrRequestLine
	SIPHdrStatusLine
	SIPHdrAccept
	SIPHdrAcceptEncoding
	SIPHdrAcceptLanguage
	SIPHdrAlertInfo
	SIPHdrAllow
	SIPHdrAuthenticationInfo
	SIPHdrAuthorization
	SIPHdrCallID
	SIPHdrCallInfo
	SIPHdrContact
	SIPHdrContentDisposition
	SIPHdrContentEncoding
	SIPHdrContentLanguage
	SIPHdrContentLength
	SIPHdrContentType
	SIPHdrCSeq
	SIPHdrDate
	SIPHdrErrorInfo
	SIPHdrExpires
	SIPHdrFrom
	SIPHdrInReplyTo
	SIPHdrMaxForwards
	SIPHdrMIMEVersion
	SIPHdrMinExpires
	SIPHdrOrganization
	SIPHdrPriority
	SIPHdrProxyAuthenticate
	SIPHdrProxyAuthorization
	SIPHdrProxyRequire
	SIPHdrRecordRoute
	SIPHdrReplyTo
	SIPHdrRequire
	SIPHdrRetryAfter
	SIPHdrRoute
	SIPHdrServer
	SIPHdrSubject
	SIPHdrSupported
	SIPHdrTimestamp
	SIPHdrTo
	SIPHdrUnsupported
	SIPHdrUserAgent
	SIPHdrVia
	SIPHdrWarning
	SIPHdrWWWAuthenticate
)

// HeadersList SIP headers list
type HeadersList []*Header

// Count number of headers
func (l HeadersList) Count() int {
	return len(l)
}

// FindByName find header by name
func (l HeadersList) FindByName(name string) *Header {
	for _, h := range l {
		if strings.EqualFold(name, h.Name()) {
			return h
		}
	}
	return nil
}

// Find find header by ID
func (l HeadersList) Find(id HdrType) *Header {
	for _, h := range l {
		if h.ID() == id {
			return h
		}
	}
	return nil
}

func (l HeadersList) exists(buf []byte) bool {
	for _, h := range l {
		if bytes.Equal(buf, h.buf) {
			return true
		}
	}
	return false
}

// Header SIP header
type Header struct {
	buf   []byte
	id    HdrType
	name  pl
	value pl
}

// ID SIP header ID
func (h *Header) ID() HdrType {
	return h.id
}

// Name SIP header name
func (h *Header) Name() string {
	return string(h.buf[h.name.p:h.name.l])
}

// Value SIP header value
func (h *Header) Value() string {
	return string(h.buf[h.value.p:h.value.l])
}

// CSeq SIP sequence number
type CSeq struct {
	Num    uint
	Method string
}

func searchParam(name string, buf []byte, params []pl) (string, bool) {
	for _, p := range params {
		prm := bytes.SplitN(buf[p.p:p.l], []byte("="), 2)
		if bytes.EqualFold([]byte(name), prm[0]) {
			if len(prm) < 2 {
				return "", true
			}
			return string(prm[1]), true
		}
	}
	return "", false
}
