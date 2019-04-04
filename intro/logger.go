package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type SqlLogger struct{}
type ConsoleLogger struct{}
type FileLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println("Console: " + message)
}

func (l SqlLogger) Log(message string) {
	fmt.Println("Sql: " + message)
}

func (l FileLogger) Log(message string) {
	fmt.Println("File: " + message)
}

func process(logger Logger) {
	logger.Log("hello!")
}

func logger() {
	process(SqlLogger{})
	process(FileLogger{})
	process(ConsoleLogger{})
}
