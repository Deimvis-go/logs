package logstest

import (
	"github.com/Deimvis-go/logs/logs"
)

var (
	FakeZap  = FakeZap1
	FakeZap1 = logs.NewZap(logs.ZapParams{
		Config: &logs.LoggerConfig{Encoding: logs.E_Console},
		Level:  logs.NewLevel(logs.LevelParams{Config: &logs.LevelConfig{Level: logs.L_Debug}}).ZapLevel,
	}).Sugar().Named("fake1")
)

func Sync() {
	FakeZap1.Sync()
}
