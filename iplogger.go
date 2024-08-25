package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
	"strings"
	"github.com/fsnotify/fsnotify"
)

const logFile = "log.txt"

func saveLog(logEntry string) {
	// Open the log file in append mode
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
	fmt.Fprintf(w, "Logged request from %s", r.RemoteAddr)
}

func startServer() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8022")
	if err := http.ListenAndServe(":8022", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func highlightDateTime(logEntry string) string {
	parts := strings.SplitN(logEntry, "] ", 2)
	if len(parts) < 2 {
		return logEntry
	}
	dateTime := parts[0] + "]"
	restOfLog := parts[1]
	highlightedDateTime := fmt.Sprintf("\033[1;37;41m%s\033[0m", dateTime)
	return fmt.Sprintf("%s %s", highlightedDateTime, restOfLog)
}

func displayLog() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating file watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(logFile)
	if err != nil {
		fmt.Println("Error adding file to watcher:", err)
		return
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				// Clear the console
				fmt.Print("\033[H\033[2J")

				// Open the log file for reading
				file, err := os.Open(logFile)
				if err != nil {
					fmt.Println("Error opening log file:", err)
					return
				}

				// Read and display the log file with highlights
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line := scanner.Text()
					fmt.Println(highlightDateTime(line))
				}
				if err := scanner.Err(); err != nil {
					fmt.Println("Error reading log file:", err)
				}
				file.Close()
			}
		case err := <-watcher.Errors:
			fmt.Println("Error watching file:", err)
		}
	}
}

func main() {
	go startServer()
	displayLog()
}
