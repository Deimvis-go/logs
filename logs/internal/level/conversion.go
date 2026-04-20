package level

import (
	"fmt"

	"github.com/Deimvis/go-ext/go1.25/ext"
)

type ConversionRules[ForeignLevel comparable] map[Level]ForeignLevel
type RevConversionRules[ForeignLevel comparable] map[ForeignLevel]Level

func ConvertTo[ForeignLevel comparable](lvl Level, r ConversionRules[ForeignLevel]) (ForeignLevel, bool) {
	flvl, ok := r[lvl]
	return flvl, ok
}

func ConvertFrom[ForeignLevel comparable](flvl ForeignLevel, r RevConversionRules[ForeignLevel]) (Level, bool) {
	lvl, ok := r[flvl]
	return lvl, ok
}

func (cr ConversionRules[ForeignLevel]) ValidateSelf() error {
	for _, lvl := range _allLevels {
		if _, ok := cr[lvl]; !ok {
			return fmt.Errorf("conversion rules are missing %s level", lvl.String())
		}
	}
	return nil
}

func (rcr RevConversionRules[ForeignLevel]) ValidateSelf() error {
	for _, lvl := range rcr {
		if !ext.Contains(_allLevels, lvl) {
			return fmt.Errorf("rev conversion rules have unknown level: %d", lvl)
		}
	}
	return nil
}
