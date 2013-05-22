package main

import (
    "net"
    "net/http"
    "io/ioutil"
    "bufio"
    "fmt"
)

func readSocket(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        fmt.Print(reader.ReadString('\n'))
    }
}

func writeSocket(conn net.Conn) {
    for {
        var message string
        _ , _ = fmt.Scanf("%s", &message)
        fmt.Fprintf(conn, message)
    }
}

func main() {
    var clientPairId string
    fmt.Print("Enter a unique ID for the client pair: ")
    _, _ = fmt.Scanf("%s", &clientPairId)
    response, _ := http.Get("http://localhost:8080/meet/" + clientPairId)
    address, _ := ioutil.ReadAll(response.Body)
    conn, _ := net.Dial("udp", string(address))
    go readSocket(conn)
    writeSocket(conn)
}
