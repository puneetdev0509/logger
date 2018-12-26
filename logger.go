package logger

import "fmt"

func Info(format string, vals ...interface{}) {
	fmt.Printf("PUNEET" + format, vals)
}
