package logger

import (
	"auth2/config"
	"bytes"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const TimeFormat = "2006-01-02 15:04:05"

var logger *logrus.Logger
var out io.Writer

func output() io.Writer {
	return out
}

func Logger() *logrus.Logger {
	return logger
}

func openLogFile(filePath, fileName string) (io.Writer, error) {
	if !strings.HasSuffix(fileName, ".log") {
		fileName = fileName + ".log"
	}
	filep := path.Join(filePath, fileName)

	file, err := os.OpenFile(filep, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

type logFormatter struct{}

func (t *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteByte('[')
	buf.WriteString(strings.ToUpper(entry.Level.String()))
	buf.WriteByte(']')
	buf.WriteByte(' ')
	buf.WriteString(entry.Time.Format(TimeFormat))
	buf.WriteByte(' ')
	if entry.HasCaller() {
		buf.WriteString(path.Base(entry.Caller.File))
		buf.WriteByte(':')
		//buf.WriteString(strings.Split(entry.Caller.Function, ".")[1])
		buf.WriteString(findFuncName(entry.Caller.Function))
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(entry.Caller.Line))
	}
	buf.WriteByte(' ')
	buf.WriteString(entry.Message)
	buf.WriteByte('\n')

	return buf.Bytes(), nil
}

func findFuncName(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return name[i+1:]
		}
	}
	return ""
}

func InitLogger() error {
	logger = logrus.New()
	logger.SetFormatter(&logFormatter{})
	logger.SetOutput(os.Stdout)
	out = os.Stdout
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)

	if config.LoggerConfig.ToFile {
		w, err := openLogFile(config.LoggerConfig.FilePath, config.LoggerConfig.FileName)
		if err != nil {
			return err
		}
		logger.SetOutput(w)
		out = w
	}

	return nil
}

// func Debug(args ...interface{}) {
// 	logger.Debug(args...)
// }

// func Info(args ...interface{}) {
// 	logger.Info(args...)
// }

// func Error(args ...interface{}) {
// 	logger.Error(args...)
// }

// func Fatal(args ...interface{}) {
// 	logger.Fatal(args...)
// }

// func Panic(args ...interface{}) {
// 	logger.Panic(args...)
// }

// func Trace(args ...interface{}) {
// 	logger.Trace(args...)
// }

// func Warn(args ...interface{}) {
// 	logger.Warn(args...)
// }

// func Warning(args ...interface{}) {
// 	logger.Warning(args...)
// }
