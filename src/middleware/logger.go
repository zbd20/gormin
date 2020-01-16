package middleware

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	filename "github.com/keepeye/logrus-filename"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/zbd20/gormin/src/config"
)

func LoggerToFile() gin.HandlerFunc {
	logConfig := config.GetConfig().Log
	logFilePath := logConfig.Path
	logFileName := logConfig.Name
	lfn := strings.Split(logFileName, ".")
	logFilePrefix := fmt.Sprintf("%v-%v", lfn[0], time.Now().Format("20060102"))

	logger := logrus.New()

	fileName := path.Join(logFilePath, fmt.Sprintf("%v.log", logFilePrefix))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logger.Errorf("open file %v error: %v", fileName, err)
		return nil
	}

	logger.Out = file

	logger.SetLevel(logrus.DebugLevel)

	// 日志rotatelogs
	logWriter, err := rotatelogs.New(
		fmt.Sprintf("%v-%v.log", logFilePrefix, "info"),
		//rotatelogs.WithLinkName(linkName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	// 日志分级，不同级别的可以分发到不同的目标
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}

	fileNameHook := filename.NewHook()
	fileNameHook.Field = "line"
	logger.AddHook(fileNameHook)

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
