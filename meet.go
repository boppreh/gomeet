package main

import (
    "net/http"
)

var singles = map[string] string {}

func meet(w http.ResponseWriter, r *http.Request) {
    if value, ok := singles[r.URL.Path]; ok {
        w.Write([]byte(value))
    } else {
        singles[r.URL.Path] = r.RemoteAddr
        w.Write([]byte(r.RemoteAddr))
    }
}

func main() {
    http.Handle("/meet/", http.StripPrefix("/meet/", http.HandlerFunc(meet)))
	http.ListenAndServe("localhost:8080", nil)
}
