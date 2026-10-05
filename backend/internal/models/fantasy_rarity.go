package models

// FantasyRarity is a card tier (Common … Mythic). Cards reference a rarity by ID,
// never by free text, so tiers can be renamed, recoloured or reordered in one place.
type FantasyRarity struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null;uniqueIndex" json:"name"`
	Slug      string `gorm:"not null;uniqueIndex" json:"slug"`
	Color     string `gorm:"not null;default:''" json:"color"` // hex colour used by the UI, e.g. #3B82F6
	SortOrder int    `gorm:"not null;default:0;index" json:"sort_order"`
}
