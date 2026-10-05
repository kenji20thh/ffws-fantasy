package models

// FantasyExchangeRequirement says a recipe needs Quantity cards of a given rarity.
// A recipe lists each rarity at most once.
type FantasyExchangeRequirement struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	RecipeID uint `gorm:"not null;uniqueIndex:idx_exchange_req" json:"recipe_id"`
	RarityID uint `gorm:"not null;uniqueIndex:idx_exchange_req;index" json:"rarity_id"`
	Quantity int  `gorm:"not null;check:chk_fantasy_exchange_requirements_qty,quantity >= 1" json:"quantity"`

	Recipe FantasyExchangeRecipe `gorm:"foreignKey:RecipeID;constraint:OnDelete:CASCADE" json:"-"`
	Rarity FantasyRarity         `gorm:"foreignKey:RarityID" json:"rarity"`
}
