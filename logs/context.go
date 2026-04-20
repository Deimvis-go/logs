package logs

import (
	"context"

	"github.com/Deimvis-go/logs/logs/internal/logctx"
)

// TODO: support options: AsUpfrontFields(), AsBackwardFields() (need to iterate over kvs and extract options), maybe use custom functions. research problem later
func CtxWith(ctx context.Context, keysAndValues ...any) context.Context {
	return logctx.WithKVs(ctx, keysAndValues...)
}
