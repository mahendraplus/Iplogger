package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
	"io"
)

const logFile = "log.txt"

func saveLog(logEntry string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(logEntry + "\n")
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)).Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s - %s - %s - %s", timestamp, r.RemoteAddr, r.Header.Get("User-Agent"), r.Method, r.URL)
	saveLog(logEntry)
	fmt.Fprintf(w, "Hello, %s", r.RemoteAddr)
}

func tailLogFile() {
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Println("Error reading log file:", err)
			return
		}
		displayLast10Logs()
	}
}

func displayLast10Logs() {
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var logs []string
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	if len(logs) > 10 {
		logs = logs[len(logs)-10:] // Get the last 10 logs
	}

	// Clear the console
	fmt.Print("\033[H\033[2J")
	fmt.Println("Last 10 Requests (Live):")
	for _, log := range logs {
		fmt.Println(log)
	}
}

func startServer() {
	http.HandleFunc("/", handler)
	fmt.Println("+-------------------------+")
	fmt.Println("|         IP-LOGGER       |")
	fmt.Println("| Created by mahendraplus |")
	fmt.Println("+-------------------------+")
	go tailLogFile()
	fmt.Println("Server started at http://localhost:8022")
	if err := http.ListenAndServe(":8022", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	startServer()
}
