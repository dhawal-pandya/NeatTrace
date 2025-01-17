# NeatTrace

**NeatTrace** is a lightweight log formatter that prettifies JSON logs in real-time. It is designed to work seamlessly with Unix pipes and is particularly useful for making logs easier to read during debugging and monitoring.

## Features

- Formats JSON logs with indentation for better readability.
- Works with standard input (`stdin`) to process logs in real-time.
- Supports both **Golang** and **Rust** implementations for flexibility.
- Can be configured to run system-wide for convenient usage with commands like `| nt`.

---

## How It Works

NeatTrace reads logs from `stdin`, formats them if they are valid JSON objects, and outputs the result to `stdout`. If the input log is not valid JSON, it is passed through unchanged.

Example:

### Input:

```json
{
  "level": "info",
  "message": "User logged in",
  "timestamp": "2024-11-29T12:34:56Z"
}
```

### Output:

```json
{
  "level": "info",
  "message": "User logged in",
  "timestamp": "2024-11-29T12:34:56Z"
}
```

---

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/dhawal-pandya/NeatTrace.git
cd NeatTrace
```

### 2. Choose an Implementation

You can build and use either the **Golang** or **Rust** version.

#### **Golang**

1. Navigate to the Golang directory:

   ```bash
   cd NeatTrace-Go
   ```

2. Build the binary:

   ```bash
   go build -o /usr/local/bin/neattrace .
   ```

3. (Optional) Create a shortcut:
   ```bash
   sudo ln -s /usr/local/bin/neattrace /usr/local/bin/nt
   ```

#### **Rust**

1. Navigate to the Rust directory:

   ```bash
   cd NeatTrace-Rust
   ```

2. Build the binary:

   ```bash
   cargo build --release
   cp target/release/neattrace /usr/local/bin/neattrace
   ```

3. (Optional) Create a shortcut:
   ```bash
   sudo ln -s /usr/local/bin/neattrace /usr/local/bin/nt
   ```

---

## Usage

### Run Manually

You can pipe any log-producing command into NeatTrace:

```bash
tail -f /var/log/system.log | nt
```

### Run System-Wide

#### **macOS** (Using `launchd`)

1. Create a `launchd` configuration file:

   ```bash
   sudo nano /Library/LaunchAgents/com.dhawal.neattrace.plist
   ```

2. Add the following content:

   ```xml
   <?xml version="1.0" encoding="UTF-8"?>
   <plist version="1.0">
   <dict>
       <key>Label</key>
       <string>com.dhawal.neattrace</string>
       <key>ProgramArguments</key>
       <array>
           <string>/usr/local/bin/neattrace</string>
       </array>
       <key>RunAtLoad</key>
       <true/>
       <key>StandardInputPath</key>
       <string>/var/log/system.log</string>
       <key>StandardOutPath</key>
       <string>/var/log/neattrace.log</string>
   </dict>
   </plist>
   ```

3. Load and start the service:

   ```bash
   sudo launchctl load -w /Library/LaunchAgents/com.dhawal.neattrace.plist
   ```

4. Verify that NeatTrace is running:
   ```bash
   launchctl list | grep com.dhawal.neattrace
   ```

#### **Debian/Ubuntu** (Using `systemd`)

1. Create a systemd service file:

   ```bash
   sudo nano /etc/systemd/system/neattrace.service
   ```

2. Add the following content:

   ```ini
   [Unit]
   Description=NeatTrace Log Formatter
   After=network.target

   [Service]
   ExecStart=/usr/local/bin/neattrace
   Restart=always
   StandardInput=file:/var/log/syslog
   StandardOutput=file:/var/log/neattrace.log

   [Install]
   WantedBy=multi-user.target
   ```

3. Enable and start the service:

   ```bash
   sudo systemctl enable neattrace
   sudo systemctl start neattrace
   ```

4. Verify that NeatTrace is running:

   ```bash
   systemctl status neattrace
   ```

5. Check logs:
   ```bash
   tail -f /var/log/neattrace.log
   ```

---

## Testing NeatTrace

Run the test for the desired implementation:

### Golang

```bash
go test ./...
```

### Rust

```bash
cargo test
```

---

## License

This project is licensed under the MIT License.
