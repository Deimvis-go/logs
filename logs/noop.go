package logs

import (
	"github.com/Deimvis-go/logs/logs/internal/noop"
	"github.com/Deimvis-go/logs/logs/internal/types"
)

var NoopKVLogger = types.NewKVLogger(noop.NewKVLoggerImpl())
