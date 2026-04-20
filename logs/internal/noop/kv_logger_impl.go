package noop

import (
	"github.com/Deimvis-go/logs/logs/internal/level"
	"github.com/Deimvis-go/logs/logs/internal/types"
)

func NewKVLoggerImpl() types.KVLoggerImpl {
	return &zapKVLoggerImpl{}
}

type zapKVLoggerImpl struct{}

func (zkvl *zapKVLoggerImpl) Log(lvl level.Level, msg string, keysAndValues ...any) {}
