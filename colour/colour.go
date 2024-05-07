// Package colour provides ANSI escape codes for colorizing console output.
package colour

// ANSI escape codes for resetting and setting text color.
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Orange = "\033[38;5;208m"
	White  = "\033[37m"
)
