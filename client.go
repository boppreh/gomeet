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
        s, _ := reader.ReadString('\n')
        fmt.Print(s)
    }
}

func writeSocket(conn net.Conn) {
    for {
        var message string
        _ , _ = fmt.Scanln("%s", &message)
        fmt.Fprintf(conn, message + "\n")
    }
}

func main() {
    var clientPairId string
    fmt.Print("Enter a unique ID for the client pair: ")
    _, _ = fmt.Scanf("%s", &clientPairId)
    response, _ := http.Get("http://localhost:8080/meet/" + clientPairId)
    address, _ := ioutil.ReadAll(response.Body)
    conn, _ := net.Dial("tcp", string(address))
    go readSocket(conn)
    writeSocket(conn)
}
