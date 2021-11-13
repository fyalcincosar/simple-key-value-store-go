package infrastructure

import (
	"io"
	"kvstore/usecases"
	"log"
	"os"
)

type Logger struct{}

func NewLogger() usecases.Logger {
	return &Logger{}
}

func (l *Logger) LogError(format string, v ...interface{}) {
	file, err := os.OpenFile("./log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}

func (l *Logger) LogHttpServer(format string, v ...interface{}) {
	file, err := os.OpenFile("./log/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}
