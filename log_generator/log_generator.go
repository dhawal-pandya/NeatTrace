package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Log represents a log entry structure
type Log struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define possible log levels and messages
	logLevels := []string{"info", "warning", "error", "debug"}
	messages := []string{
		"User logged in",
		"File not found",
		"Database connection established",
		"An error occurred",
		"File uploaded successfully",
	}

	// Generate and print random logs every 2 seconds
	for {
		log := generateRandomLog(logLevels, messages)
		printLog(log)
		time.Sleep(500 * time.Millisecond)
	}
}

// generateRandomLog creates a random log entry
func generateRandomLog(levels []string, messages []string) Log {
	level := levels[rand.Intn(len(levels))]
	message := messages[rand.Intn(len(messages))]
	timestamp := time.Now().Format(time.RFC3339)

	return Log{
		Level:     level,
		Message:   message,
		Timestamp: timestamp,
	}
}

// printLog formats and prints the log as a JSON string
func printLog(log Log) {
	// Convert the log to JSON format
	logJSON, err := json.Marshal(log)
	if err != nil {
		fmt.Printf("Error marshalling log: %v\n", err)
		return
	}

	// Print the formatted log
	fmt.Println(string(logJSON))
}
