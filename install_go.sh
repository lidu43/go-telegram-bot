#!/bin/bash
set -e

# Install Go
wget https://golang.org/dl/go1.23.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Run the build command
go mod tidy
