
# IP Logger

A simple Go application that logs HTTP requests and displays them in real-time in the CLI.

## 📥 Quick Start

1. **Download and Run the Setup Script**

   Run the following command to download and execute the setup script. This script will install Go, clone the repository, and run the Go application:

   ```bash
   curl -s https://raw.githubusercontent.com/mahendraplus/Iplogger/Max/run.sh | bash
   ```

2. **Visit the Server**

   The Go application starts a server on `http://localhost:8022`. Open this URL in your web browser to test logging.

3. **View Logs**

   Logs are displayed in real-time in the terminal where the script was executed. The log entries are saved in `log.txt`.

## 🚀 Running the Application Manually

If you prefer to run the application manually:

1. **Clone the Repository**

   ```bash
   git clone https://github.com/mahendraplus/Iplogger.git
   cd Iplogger/Max
   ```

2. **Install Dependencies**

   Make sure Go is installed and set up, then install the required package:

   ```bash
   go get github.com/fsnotify/fsnotify
   ```

3. **Run the Application**

   ```bash
   go run iplogger.go
   ```

## 📝 Notes

- **Log File**: All HTTP request logs are saved in `log.txt`.
- **Stopping the Application**: Press `Ctrl+C` in the terminal.

## 👨‍💻 Developed By

- [Mahendra Mali](https://github.com/mahendraplus)
```
