package main

import (
    "net/http"
    "fmt"
)

var singles = map[string] string {}

func meet(w http.ResponseWriter, r *http.Request) {
    if value, ok := singles[r.URL.Path]; ok {
        fmt.Println(r.RemoteAddr + " found pair at id " + r.URL.Path)
        w.Write([]byte(value))
        delete(singles, r.URL.Path)
    } else {
        fmt.Println(r.RemoteAddr + " is waiting on id " + r.URL.Path)
        singles[r.URL.Path] = r.RemoteAddr
        w.Write([]byte(r.RemoteAddr))
    }
}

func main() {
    http.Handle("/meet/", http.StripPrefix("/meet/", http.HandlerFunc(meet)))
	http.ListenAndServe("localhost:8080", nil)
}
