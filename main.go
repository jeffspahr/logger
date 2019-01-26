package main

import "time"
import logrus "github.com/sirupsen/logrus"
import "flag"

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
		time.Sleep(1 * time.Second)
	}
}
