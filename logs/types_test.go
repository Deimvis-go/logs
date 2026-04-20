package logs

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestCastToZap(t *testing.T) {
	t.Run("explicit_As", func(t *testing.T) {
		s := zap.S()
		lg, err := AsKVCtxLogger(s)
		require.NoError(t, err)
		s2, err := CastToZap(lg)
		require.NoError(t, err)
		require.NotNil(t, s2)
	})
	t.Run("implicit_As", func(t *testing.T) {
		s := zap.S()
		lg := ZapAsKVCtxLogger(s)
		s2, err := CastToZap(lg)
		require.NoError(t, err)
		require.NotNil(t, s2)
	})
}
