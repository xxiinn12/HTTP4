package client

import (
	"context"
	"crypto/tls"

	"github.com/quic-go/quic-go"
)

func Send(addr string, data []byte) ([]byte, error) {
	conn, err := quic.DialAddr(context.Background(), addr, &tls.Config{
		InsecureSkipVerify: true,
		NextProtos: []string{"h4"}, // ALPN
	}, nil)
	if err != nil {
		return nil, err
	}

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}

	stream.Write(data)

	buf := make([]byte, 4096)
	n, _ := stream.Read(buf)

	return buf[:n], nil
}
