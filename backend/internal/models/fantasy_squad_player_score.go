package models

// FantasySquadPlayerScore is the score snapshot of one card in one squad. Nothing writes to it
// yet; a later phase will fill it from real tournament statistics.
type FantasySquadPlayerScore struct {
	ID                uint `gorm:"primaryKey" json:"id"`
	SquadID           uint `gorm:"not null;uniqueIndex:idx_squad_score_card" json:"squad_id"`
	PlayerCardID      uint `gorm:"not null;uniqueIndex:idx_squad_score_card;index" json:"player_card_id"`
	BasePoints        int  `gorm:"not null;default:0" json:"base_points"`
	CaptainMultiplier int  `gorm:"not null;default:1;check:chk_fantasy_squad_player_scores_mult,captain_multiplier >= 1" json:"captain_multiplier"`
	FinalPoints       int  `gorm:"not null;default:0" json:"final_points"`

	Squad      FantasySquad      `gorm:"foreignKey:SquadID;constraint:OnDelete:CASCADE" json:"-"`
	PlayerCard FantasyPlayerCard `gorm:"foreignKey:PlayerCardID" json:"-"`
}
