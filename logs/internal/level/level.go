package level

import "fmt"

type Level uint8

var (
	L_Debug Level = 8
	L_Info  Level = 16
	L_Warn  Level = 24
	L_Error Level = 32

	_minLevel  = L_Debug
	_maxLevel  = L_Error
	_allLevels = []Level{L_Debug, L_Info, L_Warn, L_Error}
)

func (l Level) String() string {
	switch l {
	case L_Debug:
		return "debug"
	case L_Info:
		return "info"
	case L_Warn:
		return "warn"
	case L_Error:
		return "error"
	}
	panic(fmt.Errorf("got unexpected level: %d", l))
}

func (l Level) next() Level {
	if l == _maxLevel {
		return l
	}
	return l + 8
}
