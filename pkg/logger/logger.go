package logger

import (
	"com.phh/start-web/pkg/config"
	"github.com/google/wire"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var _ SimpleLogger = (*Logger)(nil)

var LoggerSet = wire.NewSet(NewLogger)

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

type SimpleLogger interface {
	Debugf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type Logger struct {
	SimpleLogger
}

func NewLogger(config *config.Config) *Logger {
	return &Logger{SimpleLogger: newZapLog(config)}
}

// getWriter
func getWriter(config *config.Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Viper.GetString("log.filename"),
		MaxSize:    config.Viper.GetInt("log.maxSize"), // megabytes
		MaxBackups: config.Viper.GetInt("log.maxBackups"),
		MaxAge:     config.Viper.GetInt("log.maxAge"), //days
		LocalTime:  config.Viper.GetBool("log.localTime"),
		Compress:   config.Viper.GetBool("log.compress"), // disabled by default
	}
}

func newZapLog(config *config.Config) *zap.SugaredLogger {
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
		return lvl >= levelMap[strings.ToLower(config.Viper.GetString("log.level"))]
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getWriter(config)))
	// 如果info、debug、error分文件记录，就创建多个 writer
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), level), // 可添加多个
	)
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return zap.New(core, zap.AddCaller()).Sugar()
}
