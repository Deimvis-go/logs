package types

import (
	"context"

	"github.com/Deimvis-go/logs/logs/internal/level"
	"github.com/Deimvis-go/logs/logs/internal/logctx"
)

func upgradeKVImplToCtx(kvImpl KVLoggerImpl) KVCtxLoggerImpl {
	return &kvCtxLoggerImpl{kvImpl: kvImpl}
}

type kvCtxLoggerImpl struct {
	kvImpl KVLoggerImpl
}

func (l *kvCtxLoggerImpl) Log(ctx context.Context, lvl level.Level, msg string, keysAndValues ...any) {
	l.kvImpl.Log(lvl, msg, l.mergeKVs(ctx, keysAndValues)...)
}

func (l *kvCtxLoggerImpl) KVImpl() KVLoggerImpl {
	return l.kvImpl
}

// TODO: implement proper merging or make sure it's not needed
func (l *kvCtxLoggerImpl) mergeKVs(ctx context.Context, keysAndValues []any) []interface{} {
	ctxKVs := logctx.KVs(ctx)
	if len(ctxKVs) == 0 {
		// fast path
		return keysAndValues
	}
	return append(ctxKVs, keysAndValues...)
}
