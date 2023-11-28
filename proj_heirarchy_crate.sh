#!/bin/bash

# Define your project directories and files
directories=(
    "api"
    "cmd/server"
    "config"
    "db"
    "db/models"
    "pkg"
    "pkg/util"
    "pkg/auth"
    "static"
    "static/images"
    "views"
)

files=(
    "api/handler.go"
    "api/middleware.go"
    "cmd/server/main.go"
    "config/config.go"
    "db/connection.go"
    "db/models/menuitem.go"
    "db/models/location.go"
    "db/models/feedback.go"
    ".env"
    ".gitignore"
    "Dockerfile"
    "go.mod"
    "go.sum"
)

# Function to create directory if it doesn't exist
create_directory() {
    if [ ! -d "$1" ]; then
        mkdir -p "$1"
        echo "Created directory: $1"
    else
        echo "Directory already exists: $1"
    fi
}

# Function to create file if it doesn't exist
create_file() {
    if [ ! -f "$1" ]; then
        touch "$1"
        echo "Created file: $1"
    else
        echo "File already exists: $1"
    fi
}

# Create directories
for dir in "${directories[@]}"; do
    create_directory "$dir"
done

# Create files
for file in "${files[@]}"; do
    create_file "$file"
done

echo "Project structure setup complete."
