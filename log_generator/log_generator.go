package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Log struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	logLevels := []string{"info", "warning", "error", "debug"}
	messages := []string{
		"User logged in",
		"File not found",
		"Database connection established",
		"An error occurred",
		"File uploaded successfully",
	}

	for {
		log := generateRandomLog(logLevels, messages)
		printLog(log)
		time.Sleep(500 * time.Millisecond)
	}
}

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

func printLog(log Log) {
	logJSON, err := json.Marshal(log)
	if err != nil {
		fmt.Printf("Error marshalling log: %v\n", err)
		return
	}

	fmt.Println(string(logJSON))
}
