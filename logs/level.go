package logs

import (
	"github.com/Deimvis/go-ext/go1.25/xcheck/xmust"
	expslog "golang.org/x/exp/slog"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// TODO: move to logscfg

type LevelParams struct {
	fx.In

	Config *LevelConfig
}

type LevelResult struct {
	fx.Out

	ZapLevel     zap.AtomicLevel
	ExpSlogLevel expslog.Level
}

func NewLevel(p LevelParams) LevelResult {
	zapLevel := xmust.Do(zap.ParseAtomicLevel(p.Config.Level))
	var expSlogLevel expslog.Level
	xmust.NoErr(expSlogLevel.UnmarshalText([]byte(p.Config.Level)))
	deps := LevelResult{
		ZapLevel:     zapLevel,
		ExpSlogLevel: expSlogLevel,
	}
	return deps
}
