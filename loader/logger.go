// Package loader
package loader

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"basic-gin-backend-module/api"
)

var (
	apiLogger *logrus.Logger
	dbLogger  *logrus.Logger
)

// loggerToFile - Define logger Filed and Output
func loggerToFile() {

	// 1. Get logger config
	logFilePath := viper.GetString("logger.path")
	// 2. Check file-path exist
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
			logrus.Fatal(err)
			panic(err)
		}
	}
	// 3. APILogger Defination
	apiLogFilePath := path.Join(logFilePath, "api")
	_, err = os.Stat(apiLogFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(apiLogFilePath, os.ModePerm); err != nil {
			logrus.Fatal(err)
			panic(err)
		}
	}
	apiLogFileName := path.Join(apiLogFilePath, "api.log")
	apiLogger = logrus.New()
	// lumberjack -> io.Writer
	apiLogger.SetOutput(&lumberjack.Logger{
		Filename:   apiLogFileName,
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})
	// logrus. InfoLevel DebugLevel
	apiLogger.SetLevel(logrus.DebugLevel)
	apiLogger.SetFormatter(&logrus.TextFormatter{})
	api.Logger = apiLogger

	// 4. dbLogger Defination
	dbLogFilePath := path.Join(logFilePath, "db")
	_, err = os.Stat(dbLogFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dbLogFilePath, os.ModePerm); err != nil {
			logrus.Fatal(err)
			panic(err)
		}
	}
	dbLogFileName := path.Join(dbLogFilePath, "db.log")
	dbLogger = logrus.New()
	// lumberjack -> io.Writer
	dbLogger.SetOutput(&lumberjack.Logger{
		Filename:   dbLogFileName,
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})
	dbLogger.SetFormatter(&logrus.TextFormatter{})
}
