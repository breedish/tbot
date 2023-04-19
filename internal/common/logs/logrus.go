package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Init() {
	logrus.SetLevel(logrus.InfoLevel)
	SetFormatter(logrus.StandardLogger())
	SetOutput(logrus.StandardLogger())
}

func SetFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if isLocalEnv, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocalEnv {
		logger.SetFormatter(&prefixed.TextFormatter{ForceFormatting: true})
	}
}

func SetOutput(logger *logrus.Logger) {
	path := "/var/log/everymeet/service/everymeet-bot-service"
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s.log", path, "%Y-%m-%d"),
		rotatelogs.WithLinkName("/var/log/everymeet/service/everymeet-bot-service.log"),
		rotatelogs.WithMaxAge(time.Second*180),
		rotatelogs.WithRotationTime(time.Second*10),
	)
	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}
	logger.SetOutput(io.MultiWriter(os.Stdout, writer))
}
