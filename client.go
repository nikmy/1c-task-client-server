package main

import (
    "bufio"
    "fmt"
    "log"
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
    if len(os.Args) < 2 {
        log.Fatal("No server IP specified")
    }

    addr := os.Args[1]
    addrData := strings.Split(os.Args[1], ":")
    if len(addrData) != 2 {
        log.Fatal("Incorrect address:port format")
    }

    ip, port := addrData[0], addrData[1]
    if net.ParseIP(ip) == nil {
        log.Fatal("Server IP is invalid")
    }

    p, parseErr := strconv.Atoi(port)
    if parseErr != nil || p < 0 || p > 65535 {
        log.Fatalf("Wrong port %d", p)
    }

    conn, dialErr := net.Dial("tcp", addr)
    if dialErr != nil {
        log.Fatal(dialErr.Error())
    }

    defer func() {
        if err := conn.Close(); err != nil {
            log.Fatal(err.Error())
        }
    }()

    if err := CLIClient(addr, conn); err != nil {
        log.Println(err.Error())
    }
}
