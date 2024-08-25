package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Log struct {
	Timestamp string `json:"timestamp"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	Method    string `json:"method"`
	Headers   string `json:"headers"`
}

func saveLog(logEntry Log) {
	// Open the log file in read-write mode
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Read existing content
	fileContent, err := os.ReadFile("log.txt")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading log file:", err)
		return
	}

	// Create new log entry
	logData := fmt.Sprintf("[%s] %s - %s - %s - %s\n", logEntry.Timestamp, logEntry.IPAddress, logEntry.UserAgent, logEntry.Method, logEntry.Headers)
	newContent := logData + string(fileContent)

	// Write new content to the file
	file.Truncate(0)   // Clear the file content
	file.Seek(0, 0)   // Move the pointer to the start of the file
	_, err = file.WriteString(newContent)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	logEntry := Log{
		Timestamp: time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)).Format("2006-01-02 15:04:05"),
		IPAddress: r.RemoteAddr,
		UserAgent: r.Header.Get("User-Agent"),
		Method:    r.Method,
		Headers:   fmt.Sprintf("%v", r.Header),
	}
	saveLog(logEntry)
	fmt.Fprintf(w, "Hello, %s!", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8022")
	http.ListenAndServe(":8022", nil)
}
