package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
    // insert prometheus library here
)

func main() {
    http.HandleFunc("/", ServeHandler)
    // insert additional handlers here
    err := http.ListenAndServe(":8083", nil)
    if err != nil {
        fmt.Println(err)
    }
}

func ServeHandler(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(1000)
    time.Sleep(time.Duration(n)*time.Millisecond)
    fmt.Fprintf(w, "Prometheus Training")
}

