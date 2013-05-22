package main

import (
    "net/http"
)

type Single struct {
    address string
    lock chan string 
}

var singles = map[string] Single {}

func meet(w http.ResponseWriter, r *http.Request) {
    if value, ok := singles[r.URL.Path]; ok {
        value.lock <- r.RemoteAddr
        w.Write([]byte(value.address))
    } else {
        single := Single{r.RemoteAddr, make(chan string)}
        singles[r.URL.Path] = single
        w.Write([]byte(<- single.lock))
        delete(singles, r.URL.Path)
    }
}

func main() {
    http.Handle("/meet/", http.StripPrefix("/meet/", http.HandlerFunc(meet)))
	//http.Handle("/", http.RedirectHandler("/view/FrontPage", http.StatusFound))
	http.ListenAndServe("localhost:8080", nil)
}
