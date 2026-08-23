package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogStats struct {
	Info  int
	Warn  int
	Error int
}

func analyzeLog(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var stats LogStats
	var errors []string
	totalLogs := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		totalLogs++

		parts := strings.SplitN(line, " ", 3)

		if len(parts) < 3 {
			continue
		}

		level := parts[1]
		message := parts[2]

		switch level {
		case "INFO":
			stats.Info++

		case "WARN":
			stats.Warn++

		case "ERROR":
			stats.Error++
			errors = append(errors, message)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("\n===== LOG ANALYZER =====")
	fmt.Printf("Total Logs : %d\n", totalLogs)
	fmt.Printf("INFO       : %d\n", stats.Info)
	fmt.Printf("WARN       : %d\n", stats.Warn)
	fmt.Printf("ERROR      : %d\n", stats.Error)

	fmt.Println("\nError Logs:")

	if len(errors) == 0 {
		fmt.Println("No errors found.")
	} else {
		for i, message := range errors {
			fmt.Printf("%d. %s\n", i+1, message)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <log-file>")
		return
	}

	filename := os.Args[1]

	err := analyzeLog(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
}