package utils

import (
	"log"
	"runtime/debug"
)

func RecoverError(err interface{}) {
	if err != nil {
		data := debug.Stack()
		log.Println(string(data))
	}
}
func CheckError(err error) bool {
	if err != nil {
		data := debug.Stack()
		log.Println(string(data))
		return true
	}
	return false
}

const (
	Debug LogLevel = 1 << iota
	Info
	Error
	Exit
)

type LogLevel uint

func LogInfo(level LogLevel, msg string) {
	data := debug.Stack()
	switch level {
	case Debug, Info, Error:
		log.Println(string(data), ":", msg)
		break
	case Exit:
		log.Fatalln(string(data), ":", msg)
		break
	}
}
