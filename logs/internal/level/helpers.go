package level

import (
	"cmp"
	"slices"

	"github.com/Deimvis/go-ext/go1.25/ext"
	"github.com/Deimvis/go-ext/go1.25/xcheck/xinvar"
	"github.com/Deimvis/go-ext/go1.25/xcheck/xmust"
	"github.com/Deimvis/go-ext/go1.25/xfn"
	"github.com/Deimvis/go-ext/go1.25/xmaps"
)

// TODO: too complex and less readable then explicit definition for RevConversionRules, so maybe not worth using it
func MakeRevConversionRules[ForeignLevel cmp.Ordered](r ConversionRules[ForeignLevel], flvls []ForeignLevel, downgradePolicy bool) RevConversionRules[ForeignLevel] {
	xmust.Eq(slices.Compare(ext.Sort(xmaps.Keys(r), xfn.Id), _allLevels), 0)
	res := make(RevConversionRules[ForeignLevel])
	cur := _minLevel // cur keeps Level so that cur.next() is an upper bound for current flvl during iteration over flvls
	ext.SortIn(&flvls, xfn.Id)
	for _, flvl := range flvls {
		if cur != _maxLevel {
			// increase cur, until we find
			for {
				nextFlvl, ok := ConvertTo(cur.next(), r)
				if ok && !(nextFlvl > flvl) {
					// condition for upperbound is not reached: nextFlvl is not upper bound for flvl
					cur = cur.next()
				} else {
					break
				}
			}
		}
		var lvl Level
		if flvl == xmust.Ok(ConvertTo(cur, r)) {
			lvl = cur
		} else {
			// resolve when flvl is between cur and cur.next()
			if downgradePolicy {
				lvl = cur
			} else {
				lvl = cur.next()
			}
		}
		xinvar.NotEq(lvl, 0)
		res[flvl] = lvl
	}
	return res
}
