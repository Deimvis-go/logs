package logs

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO: move to logscfg

type ZapParams struct {
	fx.In

	Config *LoggerConfig
	Level  zap.AtomicLevel
}

func NewZap(p ZapParams) *zap.Logger {
	presetConfig := zap.NewProductionConfig()
	if p.Config.Zap != nil {
		if p.Config.Zap.Preset != nil {
			switch *p.Config.Zap.Preset {
			case ZP_Development:
				presetConfig = zap.NewDevelopmentConfig()
			case ZP_Production:
				presetConfig = zap.NewProductionConfig()
			default:
				panic(fmt.Errorf("got unsupported zap preset: %s", *p.Config.Zap.Preset))
			}
		}
	}
	config := presetConfig
	config.Sampling = nil
	config.Level = p.Level
	config.Encoding = p.Config.Encoding
	// TODO: experiment: hide caller field for all logging (2025-01-12)
	config.EncoderConfig.CallerKey = zapcore.OmitKey
	return zap.Must(config.Build())
}
