package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/alexflint/go-arg"
)

type ResponseType string

var args struct {
	Addr string `arg:"positional,required" help:"Address to ping"`
}

func main() {
	arg.MustParse(&args)

	resp, err := http.Get(args.Addr)

	if err != nil {
		log.Fatalf("Error to connect: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read answer body: %s", err)
		return
	}

	fmt.Println(string(body))
}
