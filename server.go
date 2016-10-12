package main

import (
    "net"
    //"net/http"
    //"io/ioutil"
    "bufio"
    "fmt"
    "time"
)

func readSocket(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        s, _ := reader.ReadString('\n')
        fmt.Print(s)
    }
}

func writeSocket(conn net.Conn) {
    for {
        var message string
        _ , _ = fmt.Scanf("%s", &message)
        fmt.Fprintf(conn, message + "\n")
    }
}

func handleConnection(conn net.Conn) {
    go readSocket(conn)
    go writeSocket(conn)
}

func main() {
    var clientPairId string
    fmt.Print("Enter a unique ID for the client pair: ")
    _, _ = fmt.Scanf("%s", &clientPairId)

    firstConn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        // handle error
    }
    fmt.Fprintf(firstConn, "GET /meet/" + clientPairId + " HTTP/1.0\r\n\r\n")
    firstConn.Close()
    
    // Block to let the OS free the port.
    time.Sleep(1 * time.Millisecond)

    ln, err := net.Listen("tcp", firstConn.LocalAddr().String())
    if err != nil {
        fmt.Println("Could not open listening socket! " + err.Error())
        return
    }

    fmt.Println("Listening on address " + firstConn.LocalAddr().String())

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err.Error())
            continue
        }
        go handleConnection(conn)
    }
}
