package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // 1. Serve static files (your frontend)
    fs := http.FileServer(http.Dir("./static/"))
    http.Handle("/", fs)

    // 2. Define API Routes
    http.HandleFunc("/api/projects", getProjects)
    http.HandleFunc("/api/contact", handleContact)
	http.HandleFunc("/api/projects/all", getAllProjects)

    fmt.Println("Server starting at :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}