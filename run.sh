#!/bin/bash

# Define variables
REPO_URL="https://github.com/mahendraplus/Iplogger.git"
REPO_DIR="Iplogger/"
GO_FILE="iplogger.go"
MODULE_PATH="github.com/mahendraplus/Iplogger"

# Update and install Go if not installed
echo "Checking for Go installation..."
if ! command -v go &> /dev/null; then
    echo "Go not found. Installing Go..."
    if ! pkg install go; then
        echo "pkg install failed. Trying apt..."

        # Try installing Go with apt
        if ! apt install -y golang; then
            echo "apt install failed. Trying sudo apt..."

            # Try installing Go with sudo apt
            if ! sudo apt install -y golang; then
                echo "Failed to install Go using pkg and apt. Please install Go manually."
                exit 1
            fi
        fi
    fi
    echo "Go installed successfully."
else
    echo "Go is already installed."
fi

# Clone the repository
echo "Cloning repository..."
if [ -d "$REPO_DIR" ]; then
    echo "Repository already exists."
    cd $REPO_DIR || exit
else
    git clone $REPO_URL
    cd $REPO_DIR || exit
fi

# Initialize Go module if not present
if [ ! -f "go.mod" ]; then
    echo "Initializing Go module..."
    go mod init $MODULE_PATH
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

# Run the Go application
echo "Running the Go application..."
go run $GO_FILE
