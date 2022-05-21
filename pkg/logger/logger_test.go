package logger

import (
	"com.phh/start-web/pkg/config"
	"testing"
)

// TestZap zap 日志框架
func TestZap(t *testing.T) {
	var config = config.NewConfig("../../config")
	zapLog := newZapLog(config)
	zapLog.Debugf("debug message")
	zapLog.Infof("info message")
	zapLog.Warnf("warn message")
	zapLog.Errorf("error message:%s", "this is message")
}

func TestLogger(t *testing.T) {
	var config = config.NewConfig("../../config")
	logger := NewLogger(config)
	logger.Debugf("debug message")
	logger.Infof("info message")
	logger.Warnf("warn message")
	logger.Errorf("error message:%s", "this is message")
}
