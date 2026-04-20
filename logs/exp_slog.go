package logs

import (
	"context"
	"os"

	"github.com/Deimvis-go/logs/logs/internal/logctx"
	"go.uber.org/fx"
	expslog "golang.org/x/exp/slog"
)

type ExpSlogParams struct {
	fx.In

	Level expslog.Level
}

func NewExpSlog(p ExpSlogParams) *expslog.Logger {
	// enable context logging by default
	var handler expslog.Handler
	handler = expslog.NewJSONHandler(os.Stdout, &expslog.HandlerOptions{
		Level: p.Level,
	})
	handler = &ctxHandler{h: handler}
	logger := expslog.New(handler)
	return logger
}

type ctxHandler struct {
	// should not use ctx fields (ctxHandler impl is naive)
	h expslog.Handler
}

func (ch *ctxHandler) Enabled(ctx context.Context, lvl expslog.Level) bool {
	return ch.h.Enabled(ctx, lvl)
}

func (ch *ctxHandler) Handle(ctx context.Context, r expslog.Record) error {
	kvs := logctx.KVs(ctx)
	r.Add(kvs...)
	return ch.h.Handle(ctx, r)
}

func (ch *ctxHandler) WithAttrs(attrs []expslog.Attr) expslog.Handler {
	return ch.h.WithAttrs(attrs)
}

func (ch *ctxHandler) WithGroup(name string) expslog.Handler {
	return ch.h.WithGroup(name)
}
