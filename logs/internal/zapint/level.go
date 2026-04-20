package zapint

import (
	"github.com/Deimvis/go-ext/go1.25/xcheck/xmust"
	"github.com/Deimvis-go/logs/logs/internal/level"
	"github.com/Deimvis-go/valid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ConversionRules = level.ConversionRules[zapcore.Level]{
	level.L_Debug: zap.DebugLevel,
	level.L_Info:  zap.InfoLevel,
	level.L_Warn:  zap.WarnLevel,
	level.L_Error: zap.ErrorLevel,
}

var RevConversionRules = level.RevConversionRules[zapcore.Level]{
	zap.DebugLevel:  level.L_Debug,
	zap.InfoLevel:   level.L_Info,
	zap.WarnLevel:   level.L_Warn,
	zap.ErrorLevel:  level.L_Error,
	zap.DPanicLevel: level.L_Error,
	zap.PanicLevel:  level.L_Error,
	zap.FatalLevel:  level.L_Error,
}

func init() {
	xmust.NoErr(valid.Deep(ConversionRules))
	xmust.NoErr(valid.Deep(RevConversionRules))
}
