package main

import (
    "net/http"
    "log"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs) // Serves home.html

    log.Println("Server starting on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
