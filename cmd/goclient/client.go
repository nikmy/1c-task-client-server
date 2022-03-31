package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func CLIClient(addr string, conn net.Conn) error {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s> ", addr)
		buff, err := r.ReadBytes('\n')
		if err != nil {
			return err
		}

		if len(buff) > 4096 {
			fmt.Println("Too long message")
			continue
		}

		if string(buff) == "\\exit\n" {
			return nil
		}

		_, err = conn.Write(buff)
		if err != nil {
			return err
		}
	}
}

/**
  CLI-arguments:
      goclient <address>:<port>
*/
func main() {
	if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("\t%s <ip-address>:<port>\n", os.Args[0])
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("No server IP specified")
		return
	}

	addr := os.Args[1]
	addrData := strings.Split(os.Args[1], ":")
	if len(addrData) != 2 {
		fmt.Println("Incorrect address:port format")
		return
	}

	ip, port := addrData[0], addrData[1]
	if net.ParseIP(ip) == nil {
		fmt.Println("Server IP is invalid")
		return
	}

	p, parseErr := strconv.Atoi(port)
	if parseErr != nil || p < 0 || p > 65535 {
		fmt.Printf("Wrong port %d\n", p)
		return
	}

	conn, dialErr := net.Dial("tcp", addr)
	if dialErr != nil {
		fmt.Println(dialErr.Error())
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if err := CLIClient(addr, conn); err != nil {
		fmt.Println(err.Error())
	}
}
