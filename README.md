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
- Not yet compatible with HTTPS websites
- Requires custom server


## Advanced HTTP4 Client (Using Module)

High-performance request client using HTTP4 module (ALPN "h4").

### Usage
```bash
go run request.go host:port jumlah_request


package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xxiinn12/HTTP4/client"
)

type Result struct {
	Latency time.Duration
	Error   error
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run request.go host:port jumlah_request")
		return
	}

	addr := os.Args[1]
	total, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	results := make([]Result, total)

	var wg sync.WaitGroup
	var success int64

	startGlobal := time.Now()

	for i := 0; i < total; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			start := time.Now()

			// structured payload (HTTP-like)
			payload := []byte(fmt.Sprintf(
				"METHOD:GET\nPATH:/\nID:%d\nTIME:%d\n\n",
				id,
				time.Now().UnixNano(),
			))

			// pakai module kamu (sudah include ALPN h4)
			resp, err := client.Send(addr, payload)
			if err != nil {
				results[id] = Result{Error: err}
				return
			}

			_ = resp // bisa diparse kalau mau

			results[id] = Result{
				Latency: time.Since(start),
				Error:   nil,
			}

			atomic.AddInt64(&success, 1)

		}(i)
	}

	wg.Wait()

	totalTime := time.Since(startGlobal)

	// statistik
	var totalLatency time.Duration
	var maxLatency time.Duration

	for _, r := range results {
		if r.Error == nil {
			totalLatency += r.Latency
			if r.Latency > maxLatency {
				maxLatency = r.Latency
			}
		}
	}

	var avgLatency time.Duration
	if success > 0 {
		avgLatency = totalLatency / time.Duration(success)
	}

	fmt.Println("========== HTTP4 Benchmark ==========")
	fmt.Println("Target         :", addr)
	fmt.Println("Total Requests :", total)
	fmt.Println("Success        :", success)
	fmt.Println("Failed         :", total-int(success))
	fmt.Println("Total Time     :", totalTime)
	fmt.Println("RPS            :", float64(success)/totalTime.Seconds())
	fmt.Println("Avg Latency    :", avgLatency)
	fmt.Println("Max Latency    :", maxLatency)
}
