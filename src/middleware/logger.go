package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	log "github.com/zbd20/go-utils/blog"
	"github.com/zbd20/gormin/src/config"
)

func LoggerToFile() gin.HandlerFunc {
	logConfig := config.GetConfig().Log
	logFilePath := logConfig.Path
	logFileName := logConfig.Name

	logger := logrus.New()

	fileName := path.Join(logFilePath, logFileName)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Errorf("open file %v error: %v", fileName, err)
		return nil
	}

	logger.Out = file

	logger.SetLevel(logrus.DebugLevel)

	// 日志rotatelogs
	logWriter, err := rotatelogs.New(
		fmt.Sprintf("%v-%v.log", fileName, time.Now().Format("20060102")),
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(
		writeMap,
		&logrus.TextFormatter{},
	))

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
