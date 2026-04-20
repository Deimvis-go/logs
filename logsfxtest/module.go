package logsfxtest

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/Deimvis-go/logs/logsfx"
)

var Module = fx.Module("logs", moduleOptions...)

var moduleOptions = append(
	logsfx.ModuleOptions,
	fx.Invoke(
		func(zapLogLevel zap.AtomicLevel) {
			ZapLogLevel = zapLogLevel
		},
		func(zapSLogger *zap.SugaredLogger) {
			Zap = zapSLogger
			zap.ReplaceGlobals(zapSLogger.Desugar())
		},
	),
)
