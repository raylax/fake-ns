package core

import (
    "github.com/miekg/dns"
    "net"
    "strconv"
    "strings"
)

type handler struct {
    domainSuffix string
    handlers map[uint16]func(q dns.Question)dns.RR
}

func NewHandler(domainSuffix string) *handler {
    if !strings.HasSuffix(domainSuffix, ".") {
        domainSuffix += "."
    }
    h := &handler{
        domainSuffix: domainSuffix,
    }
    h.handlers = make(map[uint16]func(q dns.Question)dns.RR)
    h.handlers[dns.TypeA] = h.handleTypeA
    return h
}

func (handler *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
    resp := dns.Msg{}
    resp.SetReply(r)
    for _, q := range resp.Question {
        h := handler.handlers[q.Qtype]
        if h == nil {
            continue
        }
        a := h(q)
        if a == nil {
            continue
        }
        resp.Answer = append(resp.Answer, a)
    }
    w.WriteMsg(&resp)
}

// ipv4
func (handler *handler) handleTypeA(q dns.Question) dns.RR {
    prefix := GetPrefix(q.Name, handler.domainSuffix)
    if prefix == "" {
        return nil
    }
    ipStr := prefix
    switch {
    case isInt(prefix):
        i, _ := strconv.Atoi(prefix)
        ipStr = intToIpv4(i)
    case isHex(prefix):
        i, _ := strconv.ParseInt(prefix[2:], 16, 32)
        ipStr = intToIpv4(int(i))
    }
    ip := net.ParseIP(ipStr).To4()
    if ip == nil {
        return nil
    }
    return &dns.A{
        Hdr: dns.RR_Header{ Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60 },
        A:   ip,
    }
}

func isHex(str string) bool {
    if len(str) < 3 || str[0] != '0' || (str[1] != 'x' && str[1] != 'X') {
        return false
    }
    for _, c := range str[2:] {
        if (c < '0' || c > '9') && (c < 'a' || c > 'f')  && (c < 'A' || c > 'F') {
            return false
        }
    }
    return true
}

func isInt(str string) bool {
    for _, c := range str {
        if c < '0' || c > '9' {
            return false
        }
    }
    return true
}

func intToIpv4(i int) string {
    var ip string
    ip += strconv.Itoa(i >> 24 & 0xff)
    ip += "."
    ip += strconv.Itoa(i >> 16 & 0xff)
    ip += "."
    ip += strconv.Itoa(i >> 8 & 0xff)
    ip += "."
    ip += strconv.Itoa(i & 0xff)
    return ip
}