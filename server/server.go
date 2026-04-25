package server

import (
	"context"
	"crypto/tls"

	"github.com/quic-go/quic-go"
)

func Start(addr string, tlsConf *tls.Config) error {
	listener, err := quic.ListenAddr(addr, tlsConf, nil)
	if err != nil {
		return err
	}

	for {
		conn, _ := listener.Accept(context.Background())
		go handle(conn)
	}
}

func handle(conn quic.Connection) {
	for {
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			return
		}

		go func(s quic.Stream) {
			buf := make([]byte, 4096)
			n, _ := s.Read(buf)

			// echo response
			s.Write([]byte("H4 RESPONSE: " + string(buf[:n])))
			s.Close()
		}(stream)
	}
}
