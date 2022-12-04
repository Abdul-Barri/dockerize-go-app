package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "hello!")
}
func main() {
    http.HandleFunc("/hello", helloHandler)
    fmt.Println("Starting the server ................")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
