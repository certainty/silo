package ux

import "fmt"

func Info(format string, args ...any) {
	fmt.Printf(format, args...)
	fmt.Println()
}

func Error(format string, args ...any) {
	fmt.Printf(format, args...)
	fmt.Println()
}
