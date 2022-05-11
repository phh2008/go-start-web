package logger

import "testing"

// TestZap zap 日志框架
func TestZap(t *testing.T) {
	zapLog := newZapLog()
	zapLog.Debugf("debug message")
	zapLog.Infof("info message")
	zapLog.Warnf("warn message")
	zapLog.Errorf("error message:%s", "this is message")
}

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.Debugf("debug message")
	logger.Infof("info message")
	logger.Warnf("warn message")
	logger.Errorf("error message:%s", "this is message")
}
