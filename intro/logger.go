package main

import (
	"fmt"
)

// Logger interface
type Logger interface {
	Log(message string)
}

// SQLLogger sql logger
type SQLLogger struct{}

// ConsoleLogger console logger
type ConsoleLogger struct{}

// FileLogger file logger
type FileLogger struct{}

// Log console implementation of Logger
func (l ConsoleLogger) Log(message string) {
	fmt.Println("Console: " + message)
}

// Log SQL implementation of Logger
func (l SQLLogger) Log(message string) {
	fmt.Println("Sql: " + message)
}

// Log file implementation of Logger
func (l FileLogger) Log(message string) {
	fmt.Println("File: " + message)
}

func process(logger Logger) {
	logger.Log("hello!")
}

func logger() {
	process(SQLLogger{})
	process(FileLogger{})
	process(ConsoleLogger{})
}
