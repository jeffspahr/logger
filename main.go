package main

import (
	"flag"
	"time"

	logrus "github.com/sirupsen/logrus"
)

func initFlags() {
	profile := flag.String("profile", "test", "Environment profile")
	flag.Parse()

	textFormat := logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	}
	jsonFormat := logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
		},
	}

	if *profile == "dev" {
		logrus.SetFormatter(&textFormat)
	} else {
		logrus.SetFormatter(&jsonFormat)
	}
}

func main() {
	initFlags()
	for {
		logrus.Infof("I love logs")

		if time.Now().Second()%5 == 0 {
			logrus.Error("Five seconds")
		}

		if time.Now().Second()%60 == 0 {
			logrus.Warn("1 minute (╯°□°）╯︵ ┻━┻")
		}

		time.Sleep(1 * time.Second)
	}
}
