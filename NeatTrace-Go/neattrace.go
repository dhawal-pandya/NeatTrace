package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	keyColor   = "\033[34m" // Blue
	valueColor = "\033[32m" // Green
	resetColor = "\033[0m"  // Default
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
		return colorizeJSON(obj, 0)
	}
	return entry
}

func colorizeJSON(data map[string]interface{}, indentLevel int) string {
	var buffer strings.Builder
	indent := strings.Repeat("  ", indentLevel)
	buffer.WriteString("{\n")
	for key, value := range data {
		buffer.WriteString(fmt.Sprintf("%s  %s\"%s\"%s: ", indent, keyColor, key, resetColor))
		buffer.WriteString(colorizeValue(value, indentLevel+1))
		buffer.WriteString(",\n")
	}
	result := strings.TrimRight(buffer.String(), ",\n")
	result += fmt.Sprintf("\n%s}", indent)
	return result
}

func colorizeValue(value interface{}, indentLevel int) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("%s\"%s\"%s", valueColor, v, resetColor)
	case float64, int, bool, nil:
		return fmt.Sprintf("%s%v%s", valueColor, v, resetColor)
	case map[string]interface{}:
		return colorizeJSON(v, indentLevel)
	case []interface{}:
		return colorizeArray(v, indentLevel)
	default:
		return fmt.Sprintf("%s%v%s", valueColor, v, resetColor)
	}
}

func colorizeArray(arr []interface{}, indentLevel int) string {
	var buffer strings.Builder
	indent := strings.Repeat("  ", indentLevel)
	buffer.WriteString("[\n")
	for i, value := range arr {
		buffer.WriteString(fmt.Sprintf("%s  %s", indent, colorizeValue(value, indentLevel+1)))
		if i < len(arr)-1 {
			buffer.WriteString(",")
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString(fmt.Sprintf("%s]", indent))
	return buffer.String()
}
