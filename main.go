package main

import (
    "encoding/json"
    "net/http"
    "time"
)

func isEvenMinute() bool {
    currentMinute := time.Now().Minute()
    isEven := currentMinute%2 == 0
    
    if isEven {
        fmt.Printf("[DEBUG] Current minute (%d) is even. Time endpoint traffic may spike.\n", currentMinute)
    }
    
    return isEven
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    isEvenMinute()
    currentTime := time.Now().Format(time.RFC3339)
    response := map[string]string{"time": currentTime}
    
    fmt.Println("Request to /time at", currentTime)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/time", timeHandler)
    http.ListenAndServe(":8795", nil)
}
