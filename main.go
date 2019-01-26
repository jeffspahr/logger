package main

import "time"
import logrus "github.com/sirupsen/logrus"
import "flag"

func init() {
	profile := flag.String("profile", "test", "Environment profile")

	if *profile == "dev" {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			// FieldMap: FieldMap{
			// 	FieldKeyTime:  "@timestamp",
			// 	FieldKeyLevel: "@level",
			// 	FieldKeyMsg:   "@message",
			// 	FieldKeyFunc:  "@caller",
			// },
		})
	}
}

func main() {
	for {
		logrus.Infof("I love logs")
		time.Sleep(1 * time.Second)
	}
}
