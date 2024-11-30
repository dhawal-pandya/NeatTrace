package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("NeatTrace is running... Formatting logs.")
	for scanner.Scan() {
		line := scanner.Text()
		formatted := formatLog(line)
		fmt.Println(formatted)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

// formatLog formats JSON objects in log entries
func formatLog(entry string) string {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(entry), &obj); err == nil {
		return prettyPrintJSON(obj)
	}
	return entry
}

// prettyPrintJSON formats a map[string]interface{} as indented JSON
func prettyPrintJSON(data map[string]interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Sprintf("Failed to format JSON: %v", data)
	}
	return strings.TrimSpace(buffer.String())
}
