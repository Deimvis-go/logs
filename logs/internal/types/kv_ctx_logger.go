package types

import (
	"context"

	"github.com/Deimvis-go/logs/logs/internal/level"
)

func NewKVCtxLogger(impl KVLoggerImpl) KVCtxLogger {
	return &kvCtxLogger{impl: upgradeKVImplToCtx(impl)}
}

type KVCtxLogger interface {
	KVCtxLoggerImpl
	Debug(ctx context.Context, msg string, keysAndValues ...any)
	Info(ctx context.Context, msg string, keysAndValues ...any)
	Warn(ctx context.Context, msg string, keysAndValues ...any)
	Error(ctx context.Context, msg string, keysAndValues ...any)

	KV() KVLogger
}

type KVCtxLoggerImpl interface {
	Log(ctx context.Context, lvl level.Level, msg string, keysAndValues ...any)
	// TODO: WithCtxFields(ctx context.Context) KVCtxLoggerImpl
	KVImpl() KVLoggerImpl
}

type kvCtxLogger struct {
	impl KVCtxLoggerImpl
}

func (kvl *kvCtxLogger) Log(ctx context.Context, lvl level.Level, msg string, keysAndValues ...any) {
	kvl.impl.Log(ctx, lvl, msg, keysAndValues...)
}

func (kvl *kvCtxLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	kvl.impl.Log(ctx, level.L_Debug, msg, keysAndValues...)
}

func (kvl *kvCtxLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	kvl.impl.Log(ctx, level.L_Info, msg, keysAndValues...)
}

func (kvl *kvCtxLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	kvl.impl.Log(ctx, level.L_Warn, msg, keysAndValues...)
}

func (kvl *kvCtxLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	kvl.impl.Log(ctx, level.L_Error, msg, keysAndValues...)
}

func (kvl *kvCtxLogger) KV() KVLogger {
	return &kvLogger{impl: kvl.impl.KVImpl()}
}

func (kvl *kvCtxLogger) KVImpl() KVLoggerImpl {
	return kvl.impl.KVImpl()
}
