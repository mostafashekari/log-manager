Here’s the updated `README.md` with your details:

---

```markdown
# Log Manager CLI

A simple and efficient **Command Line Interface (CLI)** tool built with **Golang** for managing and filtering log files. This tool is designed to help DevOps engineers and developers quickly analyze logs, filter entries, and generate useful summaries.

---

## Features

- **Filter Logs**: Search and filter logs based on keywords.
- **Summarize Log Levels**: Generate a summary of log levels like `INFO`, `ERROR`, `WARN`, and `DEBUG`.
- **Local Processing**: Works with local log files without requiring any external dependencies.
- **Easy-to-Use**: Minimal setup with clear and straightforward commands.

---

## How It Works

### 1. **Filter Logs**
You can filter logs based on a specific keyword. For example, to find all `ERROR` logs:
```bash
go run main.go -file sample.log -filter ERROR
```

### 2. **Generate Summary**
To generate a summary of all log levels:
```bash
go run main.go -file sample.log -summary
```

### 3. **Process Without Filters**
To simply print the entire log file:
```bash
go run main.go -file sample.log
```

---

## Installation

### Prerequisites
- **Golang** installed on your system. ([Download Go](https://golang.org/dl/))
- Basic knowledge of command-line tools.

### Steps
1. Clone this repository:
   ```bash
   git clone https://github.com/mostafashekari/log-manager.git
   cd log-manager
   ```

2. Run the tool:
   ```bash
   go run main.go -file sample.log
   ```

3. Optional: Build the tool for standalone use:
   ```bash
   go build -o log-manager
   ./log-manager -file sample.log
   ```

---

## Example Log File

Here’s an example of a log file (`sample.log`) the tool can process:
```
2025-01-01 10:00:00 INFO Starting the application
2025-01-01 10:05:00 WARN Low disk space
2025-01-01 10:10:00 ERROR Unable to connect to database
2025-01-01 10:15:00 DEBUG Debugging network connection
2025-01-01 10:20:00 INFO Application running smoothly
```

---

## Command-Line Flags

| Flag         | Description                          | Example                              |
|--------------|--------------------------------------|--------------------------------------|
| `-file`      | Path to the log file (required)      | `-file sample.log`                  |
| `-filter`    | Keyword to filter log entries        | `-filter ERROR`                     |
| `-summary`   | Generate a summary of log levels     | `-summary`                          |

---

## Future Enhancements

Planned features for the next version:
- Support for JSON and CSV log formats.
- Integration with cloud services to fetch and process logs.
- Advanced filtering with regular expressions.
- Improved error handling and performance optimization.

---

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your message here"
   ```
4. Push to your branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## Contact

If you have any questions or suggestions, feel free to reach out:
- **GitHub**: [mostafashekari](https://github.com/mostafashekari)
- **Email**: mostafa.shekaari@gmail.com
```
