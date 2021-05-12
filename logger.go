package zapWrapper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

var str2level = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"error": zapcore.ErrorLevel,
}

func Init(filePath, logLevel string) {
	hook := lumberjack.Logger{
		Filename:   filePath, // 日志文件路径
		MaxSize:    128,      // megabytes
		MaxBackups: 3,        // 最多保留300个备份
		MaxAge:     7,        // days
		Compress:   true,     // 是否压缩 disabled by default
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "lineNum",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	} // 格式

	level, ok := str2level[logLevel]
	if !ok {
		level = zapcore.InfoLevel
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到控制台和文件,
		level,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	log = zap.New(core, caller, development)

	zap.ReplaceGlobals(log)
}

// NewSugar 参考了https://gist.github.com/rnyrnyrny/282fe705d6e8dc012e482582d7c8ec0b
func NewSugar(name string) *zap.SugaredLogger {
	return log.Named(name).Sugar()
}
