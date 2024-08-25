#!/bin/bash

# Define variables
REPO_URL="https://github.com/mahendraplus/Iplogger.git"
REPO_DIR="Iplogger/"
GO_FILE="iplogger.go"

# Update and install Go if not installed
echo "Checking for Go installation..."
if ! command -v go &> /dev/null; then
    echo "Go not found. Installing Go..."
    wget https://golang.org/dl/go1.20.5.linux-amd64.tar.gz -O go1.20.5.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo "Go installed successfully."
else
    echo "Go is already installed."
fi

# Clone the repository
echo "Cloning repository..."
if [ -d "$REPO_DIR" ]; then
    echo "Repository already exists."
else
    git clone $REPO_URL
    cd $REPO_DIR || exit
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

# Run the Go application
echo "Running the Go application..."
go run $GO_FILE
