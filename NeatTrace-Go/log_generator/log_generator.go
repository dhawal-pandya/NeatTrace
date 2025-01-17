package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Log struct {
	Level     string                 `json:"level"`
	SubLevels []string               `json:"sub_levels"`
	Message   string                 `json:"message"`
	Timestamp string                 `json:"timestamp"`
	Details   map[string]interface{} `json:"details"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	logLevels := []string{"info", "warning", "error", "debug"}
	subLevels := []string{"auth", "db", "api", "filesystem", "network"}
	messages := []string{
		"User logged in",
		"File not found",
		"Database connection established",
		"An error occurred",
		"File uploaded successfully",
	}

	for {
		log := generateRandomLog(logLevels, subLevels, messages)
		emitLog(log)
		time.Sleep(500 * time.Millisecond)
	}
}

func generateRandomLog(levels, subLevels, messages []string) Log {
	level := levels[rand.Intn(len(levels))]
	subLevelCount := rand.Intn(3) + 1
	subLevel := randomSubLevels(subLevels, subLevelCount)
	message := messages[rand.Intn(len(messages))]
	timestamp := time.Now().Format(time.RFC3339)
	details := generateNestedDetails()

	return Log{
		Level:     level,
		SubLevels: subLevel,
		Message:   message,
		Timestamp: timestamp,
		Details:   details,
	}
}

func randomSubLevels(subLevels []string, count int) []string {
	selected := make([]string, count)
	used := map[int]bool{}

	for i := 0; i < count; i++ {
		for {
			idx := rand.Intn(len(subLevels))
			if !used[idx] {
				selected[i] = subLevels[idx]
				used[idx] = true
				break
			}
		}
	}
	return selected
}

func generateNestedDetails() map[string]interface{} {
	return map[string]interface{}{
		"user": map[string]interface{}{
			"id":       rand.Intn(1000),
			"username": "user" + fmt.Sprint(rand.Intn(100)),
		},
		"operation": map[string]interface{}{
			"type":   "read",
			"status": "success",
			"meta": map[string]interface{}{
				"attempts": rand.Intn(3) + 1,
				"latency":  fmt.Sprintf("%dms", rand.Intn(100)),
			},
		},
		"ip_address": fmt.Sprintf("192.168.%d.%d", rand.Intn(256), rand.Intn(256)),
	}
}

func emitLog(log Log) {
	logJSON, err := json.Marshal(log)
	if err != nil {
		fmt.Printf("Error generating log: %v\n", err)
		return
	}

	fmt.Println(string(logJSON))
}
