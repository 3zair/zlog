package zlog

type Config struct {
	LogFilePath  string `yaml:"logFilePath"`
	MaxSize      int    `yaml:"maxSize"`      // 单个日志最大的文件大小. 单位: MB TODO
	MaxBackups   int    `yaml:"maxBackups"`   // 日志文件最多保存多少个备份 TODO
	MaxAge       int    `yaml:"maxAge"`       // 文件最多保存多少天 TODO
	Console      bool   `yaml:"console"`      // 是否命令行输出，开发环境可以使用
	LevelString  string `yaml:"levelString"`  // 输出的日志级别, 值：debug,info,warn,error,panic,fatal
	AppId        string `yaml:"appId"`        // 服务名 日志中心
	RotationTime int    `yaml:"rotationTime"` // 按时间间隔分割日志，单位为小时 TODO
}

func NewDefaultConfig() *Config {
	return &Config{
		MaxSize:      100,
		MaxBackups:   0,
		MaxAge:       0,
		Console:      true,
		LevelString:  "1",
		RotationTime: 24,
	}
}
