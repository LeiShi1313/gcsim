package razor

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

// When Razor's Energy is below 50%, increases Energy Recharge by 30%.
func (c *char) a4() {
	val := make([]float64, attributes.EndStatType)
	val[attributes.ER] = 0.3
	c.AddStatMod(character.StatMod{
		Base:         modifier.NewBase("er-sigil", -1),
		AffectedStat: attributes.ER,
		Amount: func() ([]float64, bool) {
			if c.Energy/c.EnergyMax < 0.5 {
				return nil, false
			}

			return val, true
		},
	})
}
