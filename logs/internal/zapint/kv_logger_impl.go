package zapint

import (
	"github.com/Deimvis/go-ext/go1.25/xcheck/xmust"
	"github.com/Deimvis-go/logs/logs/internal/level"
	"github.com/Deimvis-go/logs/logs/internal/types"
	"go.uber.org/zap"
)

func NewKVLoggerImpl(z *zap.SugaredLogger) types.KVLoggerImpl {
	return &zapKVLoggerImpl{z: z.WithOptions(zap.AddCallerSkip(3))}
}

// Locate doesn't destroy original logger, it just returns underlying *zap.SugaredLogger
func Locate(lg any) (*zap.SugaredLogger, bool) {
	ctxkv, ok := lg.(types.KVCtxLogger)
	if ok {
		lg = ctxkv.KV()
	}

	kv, ok := lg.(types.KVLogger)
	if ok {
		lg = kv.Internals().Impl()
	}

	zapkv, ok := lg.(*zapKVLoggerImpl)
	if !ok {
		return nil, false
	}
	return zapkv.z, true
}

type zapKVLoggerImpl struct {
	z *zap.SugaredLogger
}

func (zkvl *zapKVLoggerImpl) Log(lvl level.Level, msg string, keysAndValues ...any) {
	flvl := xmust.Ok(level.ConvertTo(lvl, ConversionRules))
	zkvl.z.Logw(flvl, msg, keysAndValues...)
}
