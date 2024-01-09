package logs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func Info(message string) {
	log.Info(fmt.Sprintf("LOG INFO | %v", message))
}

func Error(err error) {
	log.Error(fmt.Sprintf("LOG ERROR | %v", err.Error()))
}
