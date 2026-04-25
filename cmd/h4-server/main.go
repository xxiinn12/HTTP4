package main

import (
	"crypto/tls"
	"log"

	"github.com/you/http4/server"
)

func main() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos: []string{"h4"},
	}

	log.Fatal(server.Start(":4242", tlsConf))
}
