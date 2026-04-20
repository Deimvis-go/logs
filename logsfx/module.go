package logsfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/Deimvis-go/logs/logs"
)

// TODO: allow to configure module using logsfx.ModuleWith(...)
// options to implement:
// - configure to provide ctx kv logger and kv logger from specified one
// (e.g. logsfx.WithKVCtxLoggerFromZap)
// - configure to provide roundtripwrap with specified logger (default: any)
// (e.g. logsfx.WithRoundTripWrapFnFromKVCtxLogger)

// FX Module.
//
// Requires:
// - *LevelConfig
// - *LoggerConfig
//
// Provides:
// - *"golang.org/x/exp/slog".Logger
// - *zap.Logger
// - *zap.SugaredLogger
// - *logs.ZapCtx
// - xhttp.RoundTripWrapFn `name:"logs"`
var Module = fx.Module("logs", ModuleOptions...)

var ModuleOptions = []fx.Option{
	fx.Provide(
		fx.Private,
		logs.NewLevel,
	),
	fx.Provide(
		logs.NewExpSlog,
		fx.Annotate(
			// TODO: return as logs.KVLogger and logs.KVCtxLogger,
			// and as Self: https://pkg.go.dev/go.uber.org/fx#Self
			logs.NewZap,
			fx.OnStop(func(zapLogger *zap.Logger) error {
				// sync ignoring errors
				// https://github.com/uber-go/zap/issues/1093
				zapLogger.Sync()
				return nil
			},
			),
		),
		newZapSugared,
		logs.NewZapCtx,
		fx.Annotate(
			logs.NewZapRoundTripWrap,
			fx.ResultTags(`group:"round_trip_wraps"`),
		),
	),
}

func newZapSugared(l *zap.Logger) *zap.SugaredLogger {
	return l.Sugar()
}
