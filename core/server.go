package core

import "github.com/miekg/dns"

type server struct {
    s *dns.Server
}

func NewServer(handler dns.Handler) *server {
    s := &dns.Server{
        Addr: ":53",
        Net: "udp",
        Handler: handler,
    }
    return &server{
        s,
    }
}

func (server *server) Start() error {
    err := server.s.ListenAndServe()
    if err != nil {
        return err
    }
    return nil
}

func (server *server) Stop() {
    server.s.Shutdown()
}