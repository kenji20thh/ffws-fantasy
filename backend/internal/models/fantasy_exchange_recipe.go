package models

// FantasyExchangeRecipe defines something a team can exchange cards for. The cards required
// are listed in FantasyExchangeRequirement. No exchange logic exists yet.
type FantasyExchangeRecipe struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null;default:''" json:"description"`
	RewardType  string `gorm:"not null" json:"reward_type"` // what the reward is, e.g. "player_card"

	// RewardRarityID is the rarity of the reward when the reward is a card; NULL for other reward types.
	RewardRarityID *uint `gorm:"index" json:"reward_rarity_id"`

	// Active recipes are the only ones offered (see the note on FantasyPackType.Active about GORM defaults).
	Active bool `gorm:"not null;default:true;index" json:"active"`

	RewardRarity *FantasyRarity `gorm:"foreignKey:RewardRarityID" json:"reward_rarity,omitempty"`
}
