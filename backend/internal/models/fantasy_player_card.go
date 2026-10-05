package models

import "time"

// FantasyPlayerCard is ONE card instance owned by a fantasy team. This table IS the
// inventory: there is no separate user_inventory table.
//
// A card points at a real player but never changes it. Multiple cards for the same real
// player are allowed (even for the same team and rarity), so PlayerID is deliberately NOT unique.
type FantasyPlayerCard struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	FantasyTeamID uint      `gorm:"not null;index;index:idx_card_team_player,priority:1" json:"fantasy_team_id"`
	PlayerID      uint      `gorm:"not null;index;index:idx_card_team_player,priority:2" json:"player_id"`
	RarityID      uint      `gorm:"not null;index" json:"rarity_id"`
	Level         int       `gorm:"not null;default:1;check:chk_fantasy_player_cards_level,level >= 1" json:"level"`
	AcquiredAt    time.Time `gorm:"not null;autoCreateTime" json:"acquired_at"`

	FantasyTeam FantasyTeam   `gorm:"foreignKey:FantasyTeamID" json:"-"`
	Player      Player        `gorm:"foreignKey:PlayerID" json:"player"`
	Rarity      FantasyRarity `gorm:"foreignKey:RarityID" json:"rarity"`
}
