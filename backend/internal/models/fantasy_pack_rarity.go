package models

// FantasyPackRarity gives one rarity a weight inside one pack type. A rarity's chance is
// its weight divided by the sum of all weights of that pack, so probabilities live in the
// database rather than in code. Each (pack, rarity) pair exists at most once.
type FantasyPackRarity struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	PackTypeID uint `gorm:"not null;uniqueIndex:idx_pack_rarity" json:"pack_type_id"`
	RarityID   uint `gorm:"not null;uniqueIndex:idx_pack_rarity;index" json:"rarity_id"`
	Weight     int  `gorm:"not null;check:chk_fantasy_pack_rarities_weight,weight > 0" json:"weight"`

	PackType FantasyPackType `gorm:"foreignKey:PackTypeID;constraint:OnDelete:CASCADE" json:"-"`
	Rarity   FantasyRarity   `gorm:"foreignKey:RarityID" json:"rarity"`
}
