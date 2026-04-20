package logsfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/exp/slog"

	"github.com/Deimvis/go-ext/go1.25/xcheck/xmust"
	"github.com/Deimvis-go/logs/logs"
)

type LogLevelDependencies struct {
	fx.Out

	ZapLevel  zap.AtomicLevel
	SlogLevel slog.Level
}

func NewLogLevel(lc fx.Lifecycle, cfg *logs.Config) LogLevelDependencies {
	zapLevel := xmust.Do(zap.ParseAtomicLevel(cfg.Level))
	var slogLevel slog.Level
	xmust.NoErr(slogLevel.UnmarshalText([]byte(cfg.Level)))
	deps := LogLevelDependencies{
		ZapLevel:  zapLevel,
		SlogLevel: slogLevel,
	}
	return deps
}
