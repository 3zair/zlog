package zlog

import "github.com/sirupsen/logrus"

type ZHook struct {
	AppID string
}

func (zh *ZHook) Fire(entry *logrus.Entry) error {
	entry.Data["srv"] = zh.AppID
	return nil
}

func (zh *ZHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
