package logger

import (
	"fmt"
	"os"
	"time"
)

// LogTimestampToFile appends a timestamp and message to a specified file.
func LogTimestampToFile(message string, startTime, endTime time.Time) error {
	filename := "logtime.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	duration := endTime.Sub(startTime)

	_, err = fmt.Fprintf(file, "Start: %s: %s\n", startTime.Format("2006-01-02 15:04:05"), message)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(file, "End:   %s: %s\n", endTime.Format("2006-01-02 15:04:05"), message)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(file, "\tDuration: %s: %s\n", duration, message)
	return err
}
