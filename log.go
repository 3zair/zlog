package zlog

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

type ZLogger struct {
	Logger *logrus.Logger
}

func NewLogger() *ZLogger {
	return &zLogger
}

var zLogger ZLogger

func Init(c *Config) {
	zLogger.Logger = logrus.New()
	if c == nil {
		c = NewDefaultConfig()
	}
	var writers []io.Writer
	if c.LogFilePath != "" {
		file, err := os.OpenFile(c.LogFilePath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.Fatalf("open log file err:%v\n", err)
			return
		}
		writers = append(writers, file)
	}
	if c.Console {
		writers = append(writers, os.Stdout)
	}

	zLogger.Logger.SetOutput(io.MultiWriter(writers...))

	zLogger.Logger.SetLevel(logrusLevel[LevelStringToCode(c.LevelString)])

	if c.AppId != "" {
		zLogger.Logger.AddHook(&ZHook{
			AppID: c.AppId,
		})
	}
}

func (dl *ZLogger) InfoM(m map[string]interface{}, args ...interface{}) {
	dl.Logger.WithFields(m).Info(args...)
}
func (dl *ZLogger) ErrorM(m map[string]interface{}, args ...interface{}) {
	dl.Logger.WithFields(m).Error(args...)
}
func (dl *ZLogger) WarnM(m map[string]interface{}, args ...interface{}) {
	dl.Logger.WithFields(m).Warn(args...)
}

func (dl *ZLogger) Debugf(format string, args ...interface{}) {
	dl.Logger.Debugf(format, args...)
}
func (dl *ZLogger) Infof(format string, args ...interface{}) {
	dl.Logger.Infof(format, args...)
}
func (dl *ZLogger) Warnf(format string, args ...interface{}) {
	dl.Logger.Warnf(format, args...)
}
func (dl *ZLogger) Errorf(format string, args ...interface{}) {
	dl.Logger.Errorf(format, args...)
}
func (dl *ZLogger) Panicf(format string, args ...interface{}) {
	dl.Logger.Panicf(format, args...)
}
func (dl *ZLogger) Fatalf(format string, args ...interface{}) {
	dl.Logger.Fatalf(format, args...)
}

func (dl *ZLogger) Debug(args ...interface{}) {
	dl.Logger.Debug(args...)
}
func (dl *ZLogger) Info(args ...interface{}) {
	dl.Logger.Info(args...)
}
func (dl *ZLogger) Warn(args ...interface{}) {
	dl.Logger.Warn(args...)
}
func (dl *ZLogger) Error(args ...interface{}) {
	dl.Logger.Error(args...)
}
func (dl *ZLogger) Panic(args ...interface{}) {
	dl.Logger.Panic(args...)
}
func (dl *ZLogger) Fatal(args ...interface{}) {
	dl.Logger.Fatal(args...)
}

func InfoM(m map[string]interface{}) {
	zLogger.Logger.WithFields(m).Infoln()
}
func ErrorM(m map[string]interface{}) {
	zLogger.Logger.WithFields(m).Errorln()
}
func WarnM(m map[string]interface{}) {
	zLogger.Logger.WithFields(m).Warnln()
}

func Debugf(format string, args ...interface{}) {
	zLogger.Logger.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	zLogger.Logger.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	zLogger.Logger.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	zLogger.Logger.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	zLogger.Logger.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	zLogger.Logger.Fatalf(format, args...)
}

func Debug(args ...interface{}) {
	zLogger.Logger.Debug(args...)
}
func Info(args ...interface{}) {
	zLogger.Logger.Info(args...)
}
func Warn(args ...interface{}) {
	zLogger.Logger.Warn(args...)
}
func Error(args ...interface{}) {
	zLogger.Logger.Error(args...)
}
func Panic(args ...interface{}) {
	zLogger.Logger.Panic(args...)
}
func Fatal(args ...interface{}) {
	zLogger.Logger.Fatal(args...)
}
