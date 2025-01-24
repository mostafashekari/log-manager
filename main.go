package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// تعریف فلگ‌ها
	filePath := flag.String("file", "", "Path to the log file")
	filter := flag.String("filter", "", "Keyword to filter logs")
	summary := flag.Bool("summary", false, "Generate a summary of log levels")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a log file using the -file flag")
		os.Exit(1)
	}

	// باز کردن فایل لاگ
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	logLevels := map[string]int{"INFO": 0, "ERROR": 0, "WARN": 0, "DEBUG": 0}

	fmt.Println("Processing logs...")
	for scanner.Scan() {
		line := scanner.Text()

		// فیلتر کردن لاگ‌ها
		if *filter != "" && !strings.Contains(line, *filter) {
			continue
		}

		// شمارش سطح لاگ‌ها
		for level := range logLevels {
			if strings.Contains(line, level) {
				logLevels[level]++
			}
		}

		// نمایش خط‌های فیلتر شده
		if *filter != "" {
			fmt.Println(line)
		}
	}

	// نمایش خلاصه
	if *summary {
		fmt.Println("\nLog Summary:")
		for level, count := range logLevels {
			fmt.Printf("%s: %d\n", level, count)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
