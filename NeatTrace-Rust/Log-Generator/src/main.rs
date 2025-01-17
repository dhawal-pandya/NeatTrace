use rand::seq::SliceRandom;
use rand::{thread_rng, Rng};
use serde_json::json;
use std::{thread, time};

#[derive(Debug)]
struct Log {
    level: String,
    sub_levels: Vec<String>,
    message: String,
    timestamp: String,
    details: serde_json::Value,
}

fn main() {
    let log_levels = vec!["info", "warning", "error", "debug"];
    let sub_levels = vec!["auth", "db", "api", "filesystem", "network"];
    let messages = vec![
        "User logged in",
        "File not found",
        "Database connection established",
        "An error occurred",
        "File uploaded successfully",
    ];

    let delay = time::Duration::from_millis(500);
    let mut rng = thread_rng();

    loop {
        let log = generate_random_log(&log_levels, &sub_levels, &messages, &mut rng);
        emit_log(&log);
        thread::sleep(delay);
    }
}

fn generate_random_log(
    levels: &[&str],
    sub_levels: &[&str],
    messages: &[&str],
    rng: &mut impl Rng,
) -> Log {
    let level = levels.choose(rng).unwrap().to_string();
    let sub_level_count = rng.gen_range(1..=3);
    let sub_level = random_sub_levels(sub_levels, sub_level_count, rng);
    let message = messages.choose(rng).unwrap().to_string();
    let timestamp = chrono::Utc::now().to_rfc3339();
    let details = generate_nested_details(rng);

    Log {
        level,
        sub_levels: sub_level,
        message,
        timestamp,
        details,
    }
}

fn random_sub_levels(sub_levels: &[&str], count: usize, rng: &mut impl Rng) -> Vec<String> {
    let mut selected = vec![];
    let mut used_indices = vec![];

    while selected.len() < count {
        let idx = rng.gen_range(0..sub_levels.len());
        if !used_indices.contains(&idx) {
            selected.push(sub_levels[idx].to_string());
            used_indices.push(idx);
        }
    }
    selected
}

fn generate_nested_details(rng: &mut impl Rng) -> serde_json::Value {
    json!({
        "user": {
            "id": rng.gen_range(1..1000),
            "username": format!("user{}", rng.gen_range(1..100)),
        },
        "operation": {
            "type": "read",
            "status": "success",
            "meta": {
                "attempts": rng.gen_range(1..=3),
                "latency": format!("{}ms", rng.gen_range(1..100)),
            }
        },
        "ip_address": format!("192.168.{}.{}", rng.gen_range(0..256), rng.gen_range(0..256)),
    })
}

fn emit_log(log: &Log) {
    let log_json = serde_json::to_string(&json!({
        "level": log.level,
        "sub_levels": log.sub_levels,
        "message": log.message,
        "timestamp": log.timestamp,
        "details": log.details,
    }))
    .unwrap();

    println!("{}", log_json);
}
