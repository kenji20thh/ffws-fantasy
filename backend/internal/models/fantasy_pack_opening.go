package models

import "time"

// FantasyPackOpening records that a fantasy team opened a pack. The cards it produced are
// listed in FantasyPackOpeningCard. History is never cascaded away with its team or pack type.
type FantasyPackOpening struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	FantasyTeamID uint      `gorm:"not null;index" json:"fantasy_team_id"`
	PackTypeID    uint      `gorm:"not null;index" json:"pack_type_id"`
	OpenedAt      time.Time `gorm:"not null;autoCreateTime" json:"opened_at"`

	FantasyTeam FantasyTeam     `gorm:"foreignKey:FantasyTeamID" json:"-"`
	PackType    FantasyPackType `gorm:"foreignKey:PackTypeID" json:"pack_type"`
}
