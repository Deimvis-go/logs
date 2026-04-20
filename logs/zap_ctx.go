package logs

import (
	"context"

	"github.com/Deimvis-go/logs/logs/internal/logctx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapCtxParams struct {
	fx.In

	Sugared *zap.SugaredLogger
}

func NewZapCtx(p ZapCtxParams) *ZapCtx {
	zapctx := &ZapCtx{
		zap: p.Sugared,
	}
	return zapctx
}

// ZapCtx is the same as *zap.SugaredLogger,
// but also applies keys and values from context.
// see CtxWith for more details.
type ZapCtx struct {
	zap *zap.SugaredLogger
}

func (l *ZapCtx) DPanicw(ctx context.Context, msg string, keysAndValues ...any) {
	l.zap.DPanicw(msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) Debugw(ctx context.Context, msg string, keysAndValues ...any) {
	l.zap.Debugw(msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) Errorw(ctx context.Context, msg string, keysAndValues ...any) {
	l.zap.Errorw(msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) Fatalw(ctx context.Context, msg string, keysAndValues ...any) {
	l.zap.Fatalw(msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) Infow(ctx context.Context, msg string, keysAndValues ...any) {
	l.zap.Infow(msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) Panicw(ctx context.Context, msg string, keysAndValues ...any) {
	l.zap.Panicw(msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) Warnw(ctx context.Context, msg string, keysAndValues ...any) {
	l.zap.Warnw(msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) Logw(ctx context.Context, lvl zapcore.Level, msg string, keysAndValues ...any) {
	l.zap.Logw(lvl, msg, l.mergeKVs(ctx, keysAndValues...)...)
}

func (l *ZapCtx) ZapWithFields(ctx context.Context) *zap.SugaredLogger {
	return l.zap.With(logctx.KVs(ctx)...)
}

func (l *ZapCtx) Zap() *zap.SugaredLogger {
	return l.zap
}

// TODO: implement proper merging or make sure it's not needed
func (l *ZapCtx) mergeKVs(ctx context.Context, keysAndValues ...any) []interface{} {
	return append(logctx.KVs(ctx), keysAndValues...)
}
