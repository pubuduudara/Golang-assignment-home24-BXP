package logger

import (
	"log"
	"time"
)

// Info logs with timestamp
func Info(message string) {
	log.Printf("[INFO] %s | %s\n", time.Now().Format(time.RFC3339), message)
}

// Warn logs with timestamp
func Warn(message string) {
	log.Printf("[WARN] %s | %s\n", time.Now().Format(time.RFC3339), message)
}

// Error logs with timestamp
func Error(err error, message ...string) {
	if err != nil {
		if len(message) > 0 {
			log.Printf("[ERROR] %s | %s | %v\n", time.Now().Format(time.RFC3339), message[0], err)
		} else {
			log.Printf("[ERROR] %s | %v\n", time.Now().Format(time.RFC3339), err)
		}
	}
}
