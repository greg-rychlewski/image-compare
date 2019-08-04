package timeutil

import "time"

// Get current datetime in a specified format
func GetCurrentTime(format string) string{
    return time.Now().Format(format)
}
