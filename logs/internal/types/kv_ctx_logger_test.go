package types_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/Deimvis-go/logs/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestKVCtxLogger_CallStack(t *testing.T) {
	core, recorded := observer.New(zapcore.InfoLevel)
	z := zap.New(core, zap.AddCaller()).Sugar()

	{
		z.Info("zap log msg")
		z.Sync()

		logRecords := recorded.All()
		require.Len(t, logRecords, 1)
		log0 := logRecords[0]
		require.True(t, log0.Caller.Defined)
		require.True(t, strings.Contains(log0.Caller.String(), "kv_ctx_logger_test.go"))
	}

	lg := logs.ZapAsKVCtxLogger(z)

	{
		lg.Info(context.Background(), "kv ctx log msg")
		z.Sync()

		logRecords := recorded.All()
		require.Len(t, logRecords, 2)
		log1 := logRecords[1]
		require.True(t, log1.Caller.Defined)
		require.True(t, strings.Contains(log1.Caller.String(), "kv_ctx_logger_test.go"))
	}
}
