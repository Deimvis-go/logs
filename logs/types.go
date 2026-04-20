package logs

import (
	"errors"
	"fmt"

	"github.com/Deimvis/go-ext/go1.25/xcheck/xmust"
	"github.com/Deimvis-go/logs/logs/internal/types"
	"github.com/Deimvis-go/logs/logs/internal/zapint"
	"go.uber.org/zap"
)

type KVLogger = types.KVLogger
type KVLoggerImpl = types.KVLoggerImpl
type KVCtxLogger = types.KVCtxLogger

// TODO: add explicit ZapAsKVLogger and ZapAsKVCtxLogger in order to simplify usage in fx
// so one will be able to add logs.Module + logs.ZapAsKVCtxLogger in options and all constructors will get zap as ctx logger

func ZapAsKVCtxLogger(z *zap.SugaredLogger) KVCtxLogger {
	return xmust.Do(AsKVCtxLogger(z))
}

func ZapAsKVLogger(z *zap.SugaredLogger) KVLogger {
	return xmust.Do(AsKVLogger(z))
}

func AsKVCtxLogger(lg any) (KVCtxLogger, error) {
	kv, err := AsKVLogger(lg)
	if err != nil {
		return nil, err
	}
	return kv.KVCtx(), nil
}

func AsKVLogger(lg any) (KVLogger, error) {
	impl, err := AsKVLoggerImpl(lg)
	if err != nil {
		return nil, err
	}
	return types.NewKVLogger(impl), nil
}

// TODO: support exp/slog
func AsKVLoggerImpl(lg any) (KVLoggerImpl, error) {
	switch l := lg.(type) {
	case zap.Logger:
		lg = l.Sugar()
	case *zap.Logger:
		lg = l.Sugar()
	case zap.SugaredLogger:
		lg = &l
	}

	switch l := lg.(type) {
	case *zap.SugaredLogger:
		return zapint.NewKVLoggerImpl(l), nil

	}
	panic(fmt.Errorf("got unexpected logger type: %T", lg))
}

// CastToZap returns a view to the underlying zap logger (mutable).
// If you need an immutable copy, copy original logger first
// ( e.g. CastToZap(kvlg.Clone()) )
func CastToZap(lg any) (*zap.SugaredLogger, error) {
	z, ok := zapint.Locate(lg)
	if !ok {
		return nil, errors.New("failed to locate zap logger")
	}
	return z, nil
}

// TODO: wait if branching on type parameters will be allowed and implement func CastTo[T any](lg any) (T, error) with brancing over T
// Maybe implement with generic Locate() any, then cast result to options, and if cast is successful validate if result type equals type parameter
