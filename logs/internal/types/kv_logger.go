package types

import (
	"github.com/Deimvis-go/logs/logs/internal/level"
)

func NewKVLogger(impl KVLoggerImpl) KVLogger {
	return &kvLogger{impl: impl}
}

type KVLogger interface {
	KVLoggerImpl
	Debug(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
	Error(msg string, keysAndValues ...any)

	KVCtx() KVCtxLogger
	Internals() KVLoggerInternals
}

type KVLoggerInternals interface {
	// Impl must return type that directly implements KVLoggerImpl intertface,
	// so it can be casted to the underlying implementation type
	Impl() KVLoggerImpl
}

type KVLoggerImpl interface {
	Log(lvl level.Level, msg string, keysAndValues ...any)

	// Allow retrieving underlying logger (e.g. zap)

	// TODO: allow setting fields in Clone
	// TODO: allow updating log level in Clone
	// zap impl:
	// - check if it's increasing level: extract core, find current level
	// - if it's increasing - use logger.WithOptions(zap.IncreaseLevel)
	// - otherwise - create new zap.Core wrap that ignores underlying core filters and applies only new level filter
	//   + highlight it for users (impl example: https://github.com/uber-go/zap/issues/763#issuecomment-557733387)
	// exp_slog impl:
	// - extract current logger Handler
	// - create new Handler wrap that overrides Enabled() and applies new level filter
	// - create new logger: expslog.New(wrappedHandler)
	// Clone(opts ...CloneOption) KVLoggerImpl
}

type kvLogger struct {
	impl KVLoggerImpl
}

func (kvl *kvLogger) Log(lvl level.Level, msg string, keysAndValues ...any) {
	kvl.impl.Log(lvl, msg, keysAndValues...)
}

func (kvl *kvLogger) Debug(msg string, keysAndValues ...any) {
	kvl.impl.Log(level.L_Debug, msg, keysAndValues...)
}

func (kvl *kvLogger) Info(msg string, keysAndValues ...any) {
	kvl.impl.Log(level.L_Info, msg, keysAndValues...)
}

func (kvl *kvLogger) Warn(msg string, keysAndValues ...any) {
	kvl.impl.Log(level.L_Warn, msg, keysAndValues...)
}

func (kvl *kvLogger) Error(msg string, keysAndValues ...any) {
	kvl.impl.Log(level.L_Error, msg, keysAndValues...)
}

func (kvl *kvLogger) KVCtx() KVCtxLogger {
	return &kvCtxLogger{impl: upgradeKVImplToCtx(kvl.impl)}
}

func (kvl *kvLogger) Internals() KVLoggerInternals {
	return &kvLoggerInternals{impl: kvl.impl}
}

type kvLoggerInternals struct {
	impl KVLoggerImpl
}

func (kvli *kvLoggerInternals) Impl() KVLoggerImpl {
	return kvli.impl
}
