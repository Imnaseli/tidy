package logger

import (
	"log"
	"os"
	"time"
    "sync"
)


var (
	logFile *os.File
	once    sync.Once
)

var logFilePath = "./Logger/log.txt"

// LogToFile initializes the logger if not initialized and logs a message to the file.
func LogToFile(message string) {
	once.Do(func() {
		var err error
		logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Error opening log file:", err)
		}
	})

	if logFile != nil {
		logMessage := time.Now().Format("2006-01-02 15:04:05") + " " + message + "\n"
		_, err := logFile.WriteString(logMessage)
		if err != nil {
			log.Println("Error writing to log file:", err)
		}
	}
}
