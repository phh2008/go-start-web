package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var _ SimpleLogger = (*Logger)(nil)

type SimpleLogger interface {
	Debugf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

type Logger struct {
	SimpleLogger
}

func NewLogger() *Logger {
	return &Logger{SimpleLogger: newZapLog()}
}

// getWriter
func getWriter(fileName string) io.Writer {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		LocalTime:  true,
		Compress:   true, // disabled by default
	}
}

func newZapLog() *zap.SugaredLogger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	//encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 记录什么级别的日志
	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getWriter("./log-error.log")))
	// 如果info、debug、error分文件记录，就创建多个 writer
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), level), // 可添加多个
	)
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return zap.New(core, zap.AddCaller()).Sugar()
}
