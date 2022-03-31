package main

import (
    "flag"
    "io"
    "log"
    "net"
    "os"
    "strconv"
    "sync"
)

func serveClient(conn net.Conn, id int, f *os.File, mu *sync.Mutex) {
    defer func(conn net.Conn) {
        err := conn.Close()
        if err != nil {
            log.Println(err.Error())
        }
        log.Printf("Connection with client %d has been closed\n", id)
    }(conn)

    header := []byte(strconv.Itoa(id) + ":")
    for {
        buff := make([]byte, 4096)
        _, err := conn.Read(buff)

        switch {
        case err != nil && err != io.EOF:
            log.Println(err.Error())
            fallthrough
        case err == io.EOF:
            return
        }

        toWrite := append(header, buff...)

        mu.Lock()
        _, err = f.Write(toWrite)
        mu.Unlock()

        if err != nil {
            log.Println(err.Error())
            return
        }
    }
}

/**
  CLI-arguments:
      goserver -p <port> -f <filename>
*/
func main() {
    port := flag.Int("p", 8080, "TCP listening port")
    if *port < 0 || *port > 65535 {
        log.Fatal("Wrong port")
    }

    logfile := flag.String("f", "goserver.log", "File for logging")
    // TODO: Check file mode

    f, fErr := os.OpenFile(*logfile, os.O_CREATE|os.O_RDWR, 0777)
    if fErr != nil {
        log.Fatalf("Cannot open file: %s", *logfile)
    }

    defer func() {
        if err := f.Close(); err != nil {
            log.Println(err.Error())
        }
    }()

    conn, listenErr := net.Listen("tcp", ":"+strconv.Itoa(*port))
    if listenErr != nil {
        log.Println(listenErr.Error())
        return
    }

    nextId := 0
    var mu sync.Mutex

    for {
        client, err := conn.Accept()

        if err != nil {
            log.Println(err.Error())
            return
        }

        go serveClient(client, nextId, f, &mu)
        nextId++
    }
}
