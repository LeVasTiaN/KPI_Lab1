package main

import (
    "encoding/json"
    "net/http"
    "time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC3339)
    response := map[string]string{"time": currentTime}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/time", timeHandler)
    http.ListenAndServe(":8795", nil)
}
