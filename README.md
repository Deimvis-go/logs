# logs

logs - golang library for logging with context logging as main feature; also provides yet another generic logging interface:)

## Installation

```bash
go get github.com/Deimvis-go/logs
```

## Examples

### Context logging

```go
package main

import (
	"context"

	"github.com/Deimvis-go/logs/logs"
	"go.uber.org/zap"
)

func main() {
	z, _ := zap.NewProduction()
	lg := logs.ZapAsKVCtxLogger(z.Sugar())

	ctx := context.Background()
	ctx = logs.CtxWith(ctx, "request_id", "abc-123")

	lg.Info(ctx, "processing request", "user_id", 42)
	// Output includes: request_id=abc-123 user_id=42
}
```

### Off-the-shelf fx.Module

```go
package main

import (
	"github.com/Deimvis-go/logs/logs"
	"github.com/Deimvis-go/logs/logsfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Supply(
			&logs.LevelConfig{Level: "info"},
			&logs.LoggerConfig{Encoding: "json"},
		),
		logsfx.Module,
		fx.Invoke(func(z *zap.Logger) {
			z.Info("app started")
		}),
	).Run()
}
```
