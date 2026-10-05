package models

// FantasyPackOpeningCard links a pack opening to a card it generated. A card is minted by
// at most one opening, hence the unique PlayerCardID.
type FantasyPackOpeningCard struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	PackOpeningID uint `gorm:"not null;index" json:"pack_opening_id"`
	PlayerCardID  uint `gorm:"not null;uniqueIndex" json:"player_card_id"`

	PackOpening FantasyPackOpening `gorm:"foreignKey:PackOpeningID" json:"-"`
	PlayerCard  FantasyPlayerCard  `gorm:"foreignKey:PlayerCardID" json:"player_card"`
}
