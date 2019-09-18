%%{
    machine grammar;
    ##
    ## Grammar for SIP message.
    ## Ref. RFC3261 section 25. Augmented BNF
    ##

    CRLF            = "\r\n";
    SP              = 0x20;
    HTAB            = 0x09;
    WSP             = SP | HTAB;            # whitespace
    LWS             = ( WSP* CRLF )? WSP+;  # linear whitespace
    SWS             = LWS?;                 # sep whitespace
    HCOLON          = ( SP | HTAB )* ":" SWS;
    DQUOTE          = 0x22;                 # " double quote
    STAR            = SWS "*" SWS;          # asterisk
    EQUAL           = SWS "=" SWS;          # equal
    RAQUOT          = ">" SWS;              # right angle quote
    LAQUOT          = SWS "<";              # left angle quote
    SEMI            = SWS ";" SWS;          # semicolon
    COMMA           = SWS "," SWS;          # comma
    SLASH           = SWS "/" SWS;          # slash
    COLON           = SWS ":" SWS;          # colon

    UTF8_CONT       = 0x80..0xBF;
    UTF8_NONASCII   = ( 0xC0..0xDF UTF8_CONT{1} ) |
                      ( 0xE0..0xEF UTF8_CONT{2} ) |
                      ( 0xF0..0xF7 UTF8_CONT{3} ) |
                      ( 0xF8..0xFB UTF8_CONT{4} ) |
                      ( 0xFC..0xFD UTF8_CONT{5} );

    quoted_pair     = "\\" (0x00..0x09 | 0x0B..0x0C | 0x0E..0x7F);
    qdtext          = LWS | 0x21 | 0x23..0x5B | 0x5D..0x7E | UTF8_NONASCII;
    quoted_string   = SWS DQUOTE ( qdtext | quoted_pair )* DQUOTE;
    # TODO: pound (#) is not allowed but often used.(?)
    mark            = [\-_.!~*'()];
    unreserved      = alnum | mark;
    reserved        = [;/?:@&=+$,];
    escaped         = "%" xdigit xdigit;
    user_unreserved = [&=+$,;?/];
    param_unreserved= [\[\]/:&+$];
    paramchar       = param_unreserved | unreserved | escaped;
    pchar           = unreserved | escaped | [:@&=+$,];
    token           = ( alnum | [\-.!%*_+`'~] )+;
    word            = ( alnum | DQUOTE | [\-.!%*_+`'~()<>:\\/\[\]?{}] )+;
    Method          = "INVITE" | "ACK" | "OPTIONS" | "BYE" | "CANCEL" | "REGISTER"
                      | token;
    user            = ( unreserved | escaped | user_unreserved )+;
    password        = ( unreserved | escaped | [&=+$,] )*;
    userinfo        = user ( ":" password )? "@";
    hex4            = xdigit{1,4};
    hexseq          = hex4 ( ":" hex4 )*;
    hexpart         = hexseq | hexseq "::" hexseq? | "::" hexseq?;

    domainlabel     = alnum | alnum ( alnum | "-" )* alnum;
    toplabel        = alpha | alpha ( alnum | "-" )* alnum;
    hostname        = ( domainlabel "." )* toplabel "."?;

    IPv4address     = digit{1,3} "." digit{1,3} "." digit{1,3} "." digit{1,3};
    IPv6address     = hexpart ( ":" IPv4address )?;
    IPv6reference   = "[" IPv6address "]";

    host            = hostname | IPv4address | IPv6reference;
    port            = digit{1,5};
    hostport        = host ( ":" port )?;

    gen_value       = token | host | quoted_string;

    transport       =  "udp"i | "tcp"i | "sctp"i | "tls"i | token;
    transport_param = "transport="i transport;
    user_param      = "user="i ("phone"i | "ip"i | token);
    method_param    = "method="i Method;
    ttl_param       = "ttl="i digit{1,3};
    maddr_param     = "maddr="i host;
    lr_param        = "lr"i;
    other_param     = paramchar+ ( "=" paramchar+ )?;
    uri_parameter   = transport_param | user_param | method_param | ttl_param |
                      maddr_param | lr_param | other_param;

    hnv_unreserved  = ( param_unreserved -- "&" ) | "?";
    hnameval        = hnv_unreserved | unreserved | escaped;
    header          = hnameval+ "=" hnameval*;
    headers         = "?" header ( "&" header )*;

    generic_param   = token ( EQUAL gen_value )?;
    # in RFC3261 it is: [ [ userinfo "@" ] hostport ] 
    # but userinfo already has "@"
    srvr            = (userinfo? hostport )?;
    reg_name        = ( unreserved | escaped | [$,;:@&=+] )+;
    authority       = srvr | reg_name;
    segment         = pchar* ( ";" pchar* )*;
    path_segments   = segment ( "/" segment )*;
    abs_path        = "/" path_segments;
    net_path        = "//" authority abs_path?;

    query           = ( reserved | unreserved | escaped )*;
    hier_part       = ( net_path | abs_path ) ( "?" query )?;

    opaque_part     = ( unreserved | escaped | [\];?:@&=+$,] ) query;

    scheme_abs      = alpha ( alpha | digit | "+" | "-" | "." )*;
    scheme_sip      = "sip"i;
    scheme_sips     = "sips"i;

    SIP_URI         = ( scheme_sips | scheme_sip ) userinfo? hostport
                      ( ";" uri_parameter )* headers?;
    ABS_URI         = scheme_abs ":" ( hier_part | opaque_part );

    # status/request line
    SIP_Version     = "SIP"i "/" digit+ "." digit+;
    Reason_Phrase   = ( reserved | unreserved | escaped
                      | UTF8_NONASCII | UTF8_CONT | SP | HTAB )*;
    RequestURI      = SIP_URI | ABS_URI;

    # machines for From/To headers
    display_name    = (token LWS)* | quoted_string;

    # contact
    qvalue          = ( "0" ( "." digit{,3} )? ) | ( "1" ( "." "0"{,3} )? );
    qparam          = "q" EQUAL qvalue;
    expires_prm     = "expires" EQUAL digit+;
    
    contact_params  = qparam | expires_prm | generic_param ;
    
    # via header
    branch_cookie   = "z9hG4bK";
    via_generic     = token ( EQUAL gen_value )?;

}%%
