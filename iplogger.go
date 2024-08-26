package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const logFile = "log.txt"

func saveLog(logEntry string) {
	file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(logEntry + "\n")
}

func handler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)).Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s - %s - %s - %s", timestamp, r.RemoteAddr, r.Header.Get("User-Agent"), r.Method, r.URL)
	saveLog(logEntry)
	fmt.Fprintf(w, "Hello, %s", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8022")
	fmt.Println("To view logs in real-time, run: tail -f log.txt")
	http.ListenAndServe(":8022", nil)
}
