package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeWorkHandler)
    http.ListenAndServe(":2024", nil)
}

func homeWorkHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Привет, я домашнее задание")
}
