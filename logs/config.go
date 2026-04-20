package logs

// TODO: move to logscfg

type Config struct {
	LevelConfig  `yaml:",inline"`
	LoggerConfig `yaml:",inline"`
}

type LevelConfig struct {
	Level Level `yaml:"level" validate:"oneof=debug info warn error"`
}

type LoggerConfig struct {
	Encoding Encoding   `yaml:"encoding" validate:"oneof=console json"`
	Zap      *ZapConfig `yaml:"zap" validate:"omitnil"`
}

type ZapConfig struct {
	Preset *ZapPreset `yaml:"preset" validate:"omitnil,oneof=development production"`
}

type Level = string

var (
	L_Debug = "debug"
	L_Info  = "info"
	L_Warn  = "warn"
	L_Error = "error"
)

type Encoding = string

var (
	E_Console = "console"
	E_JSON    = "json"
)

type ZapPreset = string

var (
	ZP_Development ZapPreset = "development"
	ZP_Production  ZapPreset = "production"
)
