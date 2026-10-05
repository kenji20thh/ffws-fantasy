package models

// FantasySquadPlayer places one owned card into a slot of a squad. A card can appear only once
// per squad and a slot can hold only one card. Slots are numbered 1-4.
type FantasySquadPlayer struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	SquadID      uint `gorm:"not null;uniqueIndex:idx_squad_slot;uniqueIndex:idx_squad_card" json:"squad_id"`
	PlayerCardID uint `gorm:"not null;uniqueIndex:idx_squad_card;index" json:"player_card_id"`
	Slot         int  `gorm:"not null;uniqueIndex:idx_squad_slot;check:chk_fantasy_squad_players_slot,slot BETWEEN 1 AND 4" json:"slot"`

	Squad      FantasySquad      `gorm:"foreignKey:SquadID;constraint:OnDelete:CASCADE" json:"-"`
	PlayerCard FantasyPlayerCard `gorm:"foreignKey:PlayerCardID" json:"player_card"`
}
