package middlewares

import (
	"qgen/challange/databases"

	"github.com/gin-gonic/gin"

	"fmt"

	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func DBMiddleware(d databases.IDBAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		d.Connect()
		c.Set("db", d)
		defer d.Close()
		c.Next()
	}
}

//log to file
func LoggerToFile() gin.HandlerFunc {

	logFilePath := "/media/saber/DATA/programming/go/QGen/service/src/logs/"
	logFileName := "logs"

	//log file
	fileName := path.Join(logFilePath, logFileName)

	//write file
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//instantiate
	logger := logrus.New()

	//set output
	logger.Out = src

	//Set the log level
	logger.SetLevel(logrus.DebugLevel)

	//set rotatelogs
	logWriter, err := rotatelogs.New(
		//File name after split
		fileName+".%Y%m%d.log",

		//Generate soft link, pointing to the latest log file
		rotatelogs.WithLinkName(fileName),

		//Set the maximum storage time (7 days)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//Set the log cutting interval (1 day)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//Add Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		//Starting time
		startTime := time.Now()

		//Process the request
		c.Next()

		//End Time
		endTime := time.Now()

		//execution time
		latencyTime := endTime.Sub(startTime)

		//request method
		reqMethod := c.Request.Method

		//request routing
		reqUri := c.Request.RequestURI

		//status code
		statusCode := c.Writer.Status()

		//Request IP
		clientIP := c.ClientIP()

		//log format
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
