package logctx

import "context"

func WithKVs(ctx context.Context, keysAndValues ...any) context.Context {
	curKVs := KVs(ctx)
	if len(curKVs) == 0 {
		// fast path
		return context.WithValue(ctx, keyValuePairsCtxKey{}, keysAndValues)
	}
	// TODO: implement proper merging or make sure it's not needed
	return context.WithValue(ctx, keyValuePairsCtxKey{}, append(curKVs, keysAndValues...))
}

func KVs(ctx context.Context) []any {
	curKVs, ok := ctx.Value(keyValuePairsCtxKey{}).([]interface{})
	if !ok {
		return nil
	}
	return curKVs
}

// TODO: maybe make context key a string value and make it exportable (capitalized) in order to natively support *gin.Context.
// it sounds inappropriate to make changes to support more high-level library,
// but it will save much efforts by allowing context fields be used natively from *gin.Context
type keyValuePairsCtxKey struct{}
