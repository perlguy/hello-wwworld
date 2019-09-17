package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[1:]
    if len(name) == 0 {
        name = "W-w-world"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
