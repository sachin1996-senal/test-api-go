package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"
)

type Response struct {
    Message string `json:"message,omitempty"`
    Error   string `json:"error,omitempty"`
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")

    if strings.TrimSpace(name) == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{Error: "Invalid Input"})
        return
    }

    firstLetter := strings.ToUpper(string(name[0]))
    if firstLetter >= "A" && firstLetter <= "M" {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(Response{Message: "Hello " + name})
        return
    }

    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(Response{Error: "Invalid Input"})
}

func main() {
    log.Println("Server running on http://localhost:8080/hello-world")
    http.HandleFunc("/hello-world", handleHelloWorld)
    http.ListenAndServe(":8080", nil)
}
