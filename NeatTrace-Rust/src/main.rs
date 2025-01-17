use serde_json::Value;
use std::io::{self, BufRead};

const KEY_COLOR: &str = "\x1b[34m"; // Blue
const VALUE_COLOR: &str = "\x1b[32m"; // Green
const RESET_COLOR: &str = "\x1b[0m"; // Default

fn main() {
    let stdin = io::stdin();
    let handle = stdin.lock();

    println!("NeatTrace is running... Formatting logs.");
    for line in handle.lines() {
        match line {
            Ok(line) => {
                let formatted = format_log(&line);
                println!("{}", formatted);
            }
            Err(err) => {
                eprintln!("Error reading input: {}", err);
                std::process::exit(1);
            }
        }
    }
}

// format_log formats JSON objects in log entries
fn format_log(entry: &str) -> String {
    if let Ok(json_obj) = serde_json::from_str::<Value>(entry) {
        return colorize_json(&json_obj, 0);
    }
    entry.to_string()
}

// colorize_json formats and colorizes JSON objects recursively
fn colorize_json(data: &Value, indent_level: usize) -> String {
    let mut buffer = String::new();
    let indent = "  ".repeat(indent_level);

    if let Some(map) = data.as_object() {
        buffer.push_str("{\n");
        for (key, value) in map {
            buffer.push_str(&format!(
                "{}  {}\"{}\"{}: {}",
                indent,
                KEY_COLOR,
                key,
                RESET_COLOR,
                colorize_value(value, indent_level + 1)
            ));
            buffer.push_str(",\n");
        }
        // Remove trailing comma and newline
        if buffer.ends_with(",\n") {
            buffer.truncate(buffer.len() - 2);
        }
        buffer.push_str(&format!("\n{}}}", indent));
    }

    buffer
}

// colorize_value formats and colorizes individual JSON values
fn colorize_value(value: &Value, indent_level: usize) -> String {
    match value {
        Value::String(s) => format!("{}\"{}\"{}", VALUE_COLOR, s, RESET_COLOR),
        Value::Number(n) => format!("{}{}{}", VALUE_COLOR, n, RESET_COLOR),
        Value::Bool(b) => format!("{}{}{}", VALUE_COLOR, b, RESET_COLOR),
        Value::Null => format!("{}null{}", VALUE_COLOR, RESET_COLOR),
        Value::Object(_) => colorize_json(value, indent_level),
        Value::Array(arr) => colorize_array(arr, indent_level),
    }
}

// colorize_array formats and colorizes JSON arrays recursively
fn colorize_array(arr: &[Value], indent_level: usize) -> String {
    let mut buffer = String::new();
    let indent = "  ".repeat(indent_level);

    buffer.push_str("[\n");
    for (i, value) in arr.iter().enumerate() {
        buffer.push_str(&format!(
            "{}  {}",
            indent,
            colorize_value(value, indent_level + 1)
        ));
        if i < arr.len() - 1 {
            buffer.push_str(",");
        }
        buffer.push_str("\n");
    }
    buffer.push_str(&format!("{}]", indent));

    buffer
}
