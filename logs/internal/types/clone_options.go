package types

import "github.com/Deimvis-go/logs/logs/internal/level"

type cloneCfg struct {
	level *level.Level
}

func WithLevel(lvl level.Level) CloneOption {
	return func(cfg *cloneCfg) {
		cfg.level = &lvl
	}
}

type CloneOption func(cfg *cloneCfg)
