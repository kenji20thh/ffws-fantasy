package models

// FantasyWallet holds a fantasy team's coin balance. One wallet per fantasy team.
// Every balance change must also be written to FantasyCoinTransaction (done in a later phase).
type FantasyWallet struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	FantasyTeamID uint `gorm:"not null;uniqueIndex" json:"fantasy_team_id"`
	Coins         int  `gorm:"not null;default:0;check:chk_fantasy_wallets_coins,coins >= 0" json:"coins"`

	FantasyTeam FantasyTeam `gorm:"foreignKey:FantasyTeamID" json:"-"`
}
