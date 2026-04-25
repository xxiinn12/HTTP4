# HTTP4
H4 PROTOCOL


# HTTP4 (Experimental)

Custom protocol over QUIC using ALPN "h4".

## Features
- QUIC transport
- Custom framing
- Low latency

## Install
go mod tidy

## Run Server
go run cmd/h4-server/main.go

## Run Client
go run cmd/h4-client/main.go "hello"

## Notes
- Not compatible with HTTPS websites
- Requires custom server
