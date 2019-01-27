package main

import (
	"flag"
	"runtime"
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
			logrus.WithFields(logrus.Fields{
				"NumCPUs":   runtime.NumCPU(),
				"GoVersion": runtime.Version(),
			}).Info("5 second stats")
		}

		if time.Now().Second()%60 == 0 {
			logrus.Warn("1 minute (╯°□°）╯︵ ┻━┻")
		}
		time.Sleep(1 * time.Second)
	}
}
