package models

// FantasyPackType is a configurable pack definition (Basic, Premium, Elite, ...).
// Which rarities a pack can contain, and how likely each is, lives in FantasyPackRarity.
type FantasyPackType struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"not null;uniqueIndex" json:"slug"`
	Description string `gorm:"not null;default:''" json:"description"`
	Price       int    `gorm:"not null;default:0;check:chk_fantasy_pack_types_price,price >= 0" json:"price"` // in coins
	CardsCount  int    `gorm:"not null;default:1;check:chk_fantasy_pack_types_cards_count,cards_count >= 1" json:"cards_count"`

	// GuaranteedRarityID, when set, is the minimum rarity one card in the pack must reach.
	GuaranteedRarityID *uint `gorm:"index" json:"guaranteed_rarity_id"`

	// Active packs are the only ones offered. Note: GORM treats a Go false as "unset" on Create
	// when the column has a default, so create inactive rows with db.Select("*") or flip them afterwards.
	Active bool `gorm:"not null;default:true;index" json:"active"`

	GuaranteedRarity *FantasyRarity `gorm:"foreignKey:GuaranteedRarityID" json:"guaranteed_rarity,omitempty"`
}
