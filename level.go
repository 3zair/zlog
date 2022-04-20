package zlog

import "github.com/sirupsen/logrus"

type Level int8

const (
	DebugLevel Level = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var logrusLevel = [...]logrus.Level{
	DebugLevel: logrus.DebugLevel,
	InfoLevel:  logrus.InfoLevel,
	WarnLevel:  logrus.WarnLevel,
	ErrorLevel: logrus.ErrorLevel,
	PanicLevel: logrus.PanicLevel,
	FatalLevel: logrus.FatalLevel,
}

var levelLowerNames = [...]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warn",
	ErrorLevel: "error",
	PanicLevel: "panic",
	FatalLevel: "fatal",
}

func (l Level) ToLowerString() string {
	return levelLowerNames[l]
}

func LevelStringToCode(levelString string) Level {
	for i, ls := range levelLowerNames {
		if ls == levelString {
			return Level(i)
		}
	}
	return DebugLevel
}
