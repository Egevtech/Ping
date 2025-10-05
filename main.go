package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ResponseType string

const (
	HTTPS ResponseType = "https"
	HTTP  ResponseType = "http"
)

func main() {

	var addr string
	var RType ResponseType
	var commands []string

	addr = "localhost:8080"
	RType = HTTP

	fmt.Printf("|rtype <type>| where type=={https, http} - select response type\n" +
		"|connect <addr>| - set address, (for ex. 'connect localhost:8080')\n")

	for {
		fmt.Printf("<%s://%s/> ?", RType, addr)
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = command[:len(command)-1]

		commands = strings.Split(command, " ")

		switch commands[0] {
		case "exit":
			return

		case "connect":
			if len(commands) < 2 {
				fmt.Printf("[WARNING] Not enough arguments: given %d, needed %d\n", len(commands), 2)
				break
			}

			addr = commands[1]
			break

		case "get":
			resp, err := http.Get(fmt.Sprintf("%s://%s/", RType, addr))
			if err != nil {
				fmt.Printf("[NF] Error occured: %s\n", err)
				break
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("[NF] Error occured while reading answer: %s\n", err)
			}
			fmt.Println(string(body))
			break

		case "rtype":
			if len(commands) < 2 {
				fmt.Printf("[WARNING] Not enough arguments: given %d, needed %d\n", len(commands), 2)
				break
			}

			switch commands[1] {
			case "https":
				RType = HTTPS
				break
			case "http":
				RType = HTTP
				break
			}

		default:
			break
		}
	}

}
