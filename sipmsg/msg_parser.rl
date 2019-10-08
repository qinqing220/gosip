// -*-go-*-
//
// SIP headers and first line parser
package sipmsg

import (
    "bytes"
)

var ErrorSIPHeader = errorNew("Invalid SIP Header")

%% machine msg;
%% write data;

func parseHeader(msg *Message, data []byte) (HdrType, error) {
    cs := 0 // current state. entery point = 0
    l := ptr(len(data))
    pos := make([]pl, 0, 12)
    params := make([]pl, 0, 12)
    var p, // data pointer
        m, // marker
        pe ptr = 0, 0, l
    var dname,         // display name
        trans,
        addr,
        port,
        ttl,
        maddr,
        recvd,
        branch,
        tag pl;        // to/from tag

    hidx := 0 // header value index

    var id HdrType

    if bytes.Equal(data, []byte("\r\n")) {
        return MsgEOF, nil
    }
%%{

    action sm        { m = p }
    action push      { pos = append(pos, pl{m, p}) }
    action tag       { tag = pl{m, p} }
    action dname     { dname.p =m; dname.l = p }
    action addr      { addr.p = m; addr.l = p }
    action port      { port.p = m; port.l = p }
    action trans     { trans.p = m; trans.l = p }
    action param     { params = append(params, pl{m, p}) }
    action reset_cnt { hidx = msg.Contacts.Count(); params = make([]pl, 0, 12) }
    action init_via  { hidx = msg.Vias.Count() }
    action reset_via {
        branch.p = 0; branch.l = 0
        ttl.p    = 0; ttl.l    = 0
        maddr.p  = 0; maddr.l  = 0
        recvd.p  = 0; recvd.l  = 0
    }
    action contact   { msg.setContact(data[:], pos[0], dname, addr, params, hidx) }
    action via       {
        msg.setVia(data[:], pos[0], trans, addr, port, branch, ttl, maddr, recvd, hidx)
    }
    action reset_route { params = make([]pl, 0, 12) }
    action route     { msg.setRoute(id, data[:], pos[0], dname, addr, params) }

    include grammar "grammar.rl";

    # -- COMMA decreases machines but fails to parse , in To header username
    addr_spec       = (SIP_URI | ABS_URI) >sm %addr;
    tag_param       = "tag"i EQUAL token >sm %tag;
    fromto_gparam   = (token -- "tag"i) >sm ( EQUAL gen_value )? %param;
    name_addr       = (display_name >sm %dname)? LAQUOT addr_spec RAQUOT;
    param_tofrom    = tag_param | fromto_gparam;
    tofrom_value    = ( name_addr | (addr_spec -- SEMI) ) ( SEMI param_tofrom )*;
    contact_value   = (( name_addr | (addr_spec -- SEMI)) ( SEMI contact_params >sm %param )* )
                      >reset_cnt %contact;

    via_ttl         = "ttl"i EQUAL digit{1,3} >sm %{ ttl.p = m; ttl.l = p };
    via_maddr       = "maddr"i EQUAL host >sm %{ maddr.p = m; maddr.l = p };
    via_received    = "received"i EQUAL (IPv4address | IPv6address) >sm %{ recvd.p = m; recvd.l = p};
    via_branch      = "branch"i EQUAL (branch_cookie token) >sm %{ branch.p = m; branch.l = p };
    via_params      = via_ttl | via_maddr | via_received | via_branch | via_generic;
    via_sent_proto  = "SIP" SLASH digit "." digit SLASH >init_via transport >sm %trans;
    sent_by         = host >sm %addr (COLON port >sm %port)?;
    via_parm        = ( via_sent_proto LWS sent_by (SEMI via_params)* )
                      >reset_via %via;
    route_param     = ( name_addr ( SEMI generic_param >sm %param )* ) >reset_route %route;

    # @Status-Line@
    StatusLine  = SIP_Version >sm %push SP digit{3} >sm %push SP
                  Reason_Phrase >sm %push CRLF @{ id = msg.setStatusLine(data, pos) };
    # @Request-Line@
    RequestLine = Method >sm %push SP RequestURI >sm %push SP
                  SIP_Version >sm %push CRLF @{ id = msg.setRequestLine(data, pos) };
    # @CSeq@
    CSeq        = name_cseq >sm %push HCOLON digit{1,10} >sm %push
                  LWS Method >sm %push CRLF @{ id = msg.setCSeq(data, pos) };
    # @Call-ID@
    CallID      = name_callid >sm %push HCOLON
                  ( word ( "@" word )? ) >sm %push CRLF @{ id = msg.setCallID(data, pos) };
    # @Content-Length@
    ContentLen  = name_cnt_len >sm %push HCOLON
                  digit{1,10} >sm %push CRLF @{ id = msg.setContentLen(data, pos) };
    # @From@
    From        = name_from >sm %push HCOLON tofrom_value CRLF
                  @{ id = msg.setFrom(data, params, pos[0], dname, addr, tag) };
    # @To@
    To          = name_to >sm %push HCOLON tofrom_value CRLF
                  @{ id = msg.setTo(data, params, pos[0], dname, addr, tag) };
    # @Contact@
    Contact     = name_contact >sm %push HCOLON
                  ( STAR %{ msg.setContactStar() } | 
                  ( contact_value ( COMMA contact_value )* )) CRLF
                  @{ id = SIPHdrContact; };
    # @Via@
    Via         = name_via >sm %push HCOLON via_parm
                  ( COMMA via_parm )* CRLF @{ id = SIPHdrVia }; 
    # @Route@
    Route       = name_route >sm %push HCOLON %{id = SIPHdrRoute}
                  route_param (COMMA route_param)* CRLF;
    # @Record-Route@
    RecordRoute = name_rroute >sm %push HCOLON %{id = SIPHdrRecordRoute}
                  route_param (COMMA route_param)* CRLF;
    # @Expires@
    Expires     = name_expires >sm %push HCOLON digit{1,10} >sm 
                  %{ id = msg.setExpires(data[m:p]) } CRLF
                  @{ msg.pushHeader(SIPHdrExpires, data, pos[0], pl{m, p}) };
    # @Max-Forwards@
    MaxForwards = name_maxfwd >sm %push HCOLON digit{1,6} >sm 
                  %{ id = msg.setMaxFwd(data[m:p]) } CRLF
                  @{ msg.pushHeader(SIPHdrMaxForwards, data, pos[0], pl{m, p}) };
    # Generic headers
    Accept      = name_accept >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrAccept) };
    AcceptEnc   = name_acc_enc >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrAcceptEncoding) };
    AcceptLang  = name_acc_lang >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrAcceptLanguage) };
    AlertInfo   = name_alert >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrAlertInfo) };
    Allow       = name_allow >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrAllow) };
    AuthInfo    = name_auth_info >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrAuthenticationInfo) };
    Auth        = name_auth >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrAuthorization) };
    CallInfo    = name_call_info >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrCallInfo) };
    ContDispo   = name_cont_disp >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrContentDisposition) };
    ContEncode  = name_cont_enc >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrContentEncoding) };
    ContLang    = name_cont_lang >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrContentLanguage) };
    ContType    = name_cont_type >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrContentType) };
    Date        = name_date >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrDate) };
    ErrorInfo   = name_err_info >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrErrorInfo) };
    InReplyTo   = name_in_reply >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrInReplyTo) };
    MIMEVer     = name_mime_ver >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrMIMEVersion) };
    MinExpires  = name_min_expr >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrMinExpires) };
    Organization= name_organizn >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrOrganization) };
    Priority    = name_priority >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrPriority) };
    PxyAuthen   = name_pauthen >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrProxyAuthenticate) };
    PxyAuthor   = name_pauthor >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrProxyAuthorization) };
    PxyRequired = name_prequired >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrProxyRequire) };
    ReplyTo     = name_reply_to >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrReplyTo) };
    Require     = name_require >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrRequire) };
    RetryAfter  = name_retryafter >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrRetryAfter) };
    Server      = name_server >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrServer) };
    Subject     = name_subject >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrSubject) };
    Supported   = name_supported >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrSupported) };
    Timestamp   = name_timestamp >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrTimestamp) };
    Unsupported = name_unsupport >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrUnsupported) };
    UserAgent   = name_user_agent >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrUserAgent) };
    Warning     = name_warning >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrWarning) };
    WWWAuth     = name_www_auth >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrWWWAuthenticate) };

    GenericHeader = header_name >sm %push HCOLON %sm header_value %push CRLF
                  @{ id = msg.setGenericHeader(data, pos, SIPHdrGeneric) };

    siphdr :=   StatusLine
              | RequestLine
              | CSeq
              | CallID
              | Contact
              | ContentLen
              | Expires
              | From
              | MaxForwards
              | RecordRoute
              | Route
              | To
              | Via
              | Accept
              | AcceptEnc
              | AcceptLang
              | AlertInfo
              | Allow
              | AuthInfo
              | Auth
              | CallInfo
              | ContDispo
              | ContEncode
              | ContLang
              | ContType
              | Date
              | ErrorInfo
              | InReplyTo
              | MIMEVer
              | MinExpires
              | Organization
              | Priority
              | PxyAuthen
              | PxyAuthor
              | PxyRequired
              | ReplyTo
              | Require
              | RetryAfter
              | Server
              | Subject
              | Supported
              | Timestamp
              | Unsupported
              | UserAgent
              | Warning
              | WWWAuth
              | GenericHeader;
}%%
    %% write init;
    %% write exec;
    if cs >= msg_first_final {
        return id, nil
    }
    return -1, ErrorSIPHeader.msg("%s", data)
}
