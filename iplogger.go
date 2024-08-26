package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const logFile = "log.txt"

func saveLog(logEntry string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating log file:", err)
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

func showLogo() {
	fmt.Println(`
+-------------------------+
|         IP-LOGGER       |
| Created by mahendraplus |
+-------------------------+`)
}

func promptLiveLog() {
	fmt.Print("Do you want to see the live log? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response) // Properly trim newline and spaces

	if response == "y" || response == "Y" {
		cmd := exec.Command("tail", "-f", "$HOME/Iplogger/log.txt")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running tail command:", err)
		}
	} else {
		fmt.Println("Live log viewing skipped. Server is running...")
	}
}

func main() {
	showLogo()
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8022")
	go promptLiveLog() // Ask user if they want to see live log
	if err := http.ListenAndServe(":8022", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
