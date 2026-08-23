# 📊 CLI Log Analyzer

A simple command-line log analyzer built with Go.

The project reads a log file line-by-line, parses log levels, calculates
statistics, and extracts error messages.

## 🚀 Features

- Read log files
- Parse log entries
- Count `INFO`, `WARN`, and `ERROR` logs
- Extract error messages
- Command-line file input
- Basic error handling
- Skip empty lines

## 🛠️ Tech Stack

- Go (Golang)
- Go Standard Library

## 📂 Project Structure

```text
03-log-analyzer/
├── go.mod
├── main.go
├── sample.log
└── README.md