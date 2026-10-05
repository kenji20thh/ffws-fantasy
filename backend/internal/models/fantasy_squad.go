package models

import "time"

// FantasySquad is a fantasy team's lineup for one tournament day. It is built from cards the
// team owns (FantasySquadPlayer), never from the global player pool. The "exactly 4 cards"
// rule is enforced by the squad feature itself in a later phase, not by this table.
type FantasySquad struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	FantasyTeamID   uint `gorm:"not null;uniqueIndex:idx_squad_team_day" json:"fantasy_team_id"`
	TournamentDayID uint `gorm:"not null;uniqueIndex:idx_squad_team_day;index" json:"tournament_day_id"`

	// CaptainCardID is NULL while the squad is still a draft with no captain chosen.
	CaptainCardID *uint `gorm:"index" json:"captain_card_id"`

	CreatedAt time.Time  `json:"created_at"`
	LockedAt  *time.Time `json:"locked_at"`

	FantasyTeam   FantasyTeam        `gorm:"foreignKey:FantasyTeamID" json:"-"`
	TournamentDay TournamentDay      `gorm:"foreignKey:TournamentDayID" json:"-"`
	CaptainCard   *FantasyPlayerCard `gorm:"foreignKey:CaptainCardID" json:"captain_card,omitempty"`
}
