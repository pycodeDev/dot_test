package helper

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type myMap map[string]interface{}

type out struct {
	Func   string
	Script string
}

type tes struct {
	Data    out
	Level   string
	Message string
}

func init() {
	var filename string = "log/error.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.JSONFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "2006-01-02 15:04:05.0000"
	Formatter.PrettyPrint = false
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
}

func LogError(errorMsg string, tags ...string) {
	parsed := pasrseField(tags...)
	log.WithFields(parsed).Error(errorMsg)
	data := make(map[string]interface{})
	data["data"] = parsed
	data["message"] = errorMsg
	data["level"] = "error"
}

func pasrseField(tags ...string) log.Fields {
	result := make(log.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}
