package main

import (
    "fake-ns/core"
    "flag"
    "os"
)

var domain string

func main() {
    flag.StringVar(&domain, "d", "", "ns domain")
    flag.Parse()
    if domain == "" {
        flag.Usage()
        os.Exit(0)
    }
    s := core.NewServer(core.NewHandler(domain))
    err := s.Start()
    defer s.Stop()
    if err != nil {
        panic(err)
    }
}
