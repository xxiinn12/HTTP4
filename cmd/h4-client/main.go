package main

import (
	"fmt"
	"os"

	"github.com/you/http4/client"
)

func main() {
	resp, err := client.Send("localhost:4242", []byte(os.Args[1]))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp))
}
