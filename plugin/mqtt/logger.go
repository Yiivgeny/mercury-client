package mqtt

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log func(msg string, fields ...zap.Field)

var LevelEnablerNone = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
	return false
})

func NewZapWrapper(logger *zap.Logger, level zapcore.Level) *ZapWrapper {
	internal := logger.WithOptions(zap.AddCallerSkip(1))

	var call log
	switch level {
	case zapcore.DebugLevel:
		call = internal.Debug
	case zapcore.InfoLevel:
		call = internal.Info
	case zapcore.WarnLevel:
		call = internal.Warn
	case zapcore.ErrorLevel:
		call = internal.Error
	case zapcore.DPanicLevel:
		call = internal.DPanic
	case zapcore.PanicLevel:
		call = internal.Panic
	case zapcore.FatalLevel:
		call = internal.Fatal
	default:
		call = internal.Debug
	}
	return &ZapWrapper{
		call: call,
	}
}

type ZapWrapper struct {
	call log
}

func (r *ZapWrapper) Println(v ...interface{}) {
	r.call(fmt.Sprint(v...))
}
func (r *ZapWrapper) Printf(format string, v ...interface{}) {
	r.call(fmt.Sprintf(format, v...))
}
