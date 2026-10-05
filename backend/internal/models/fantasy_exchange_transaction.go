package models

import "time"

// FantasyExchangeTransaction records that a fantasy team performed an exchange recipe.
type FantasyExchangeTransaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	FantasyTeamID uint      `gorm:"not null;index" json:"fantasy_team_id"`
	RecipeID      uint      `gorm:"not null;index" json:"recipe_id"`
	CreatedAt     time.Time `json:"created_at"`

	FantasyTeam FantasyTeam           `gorm:"foreignKey:FantasyTeamID" json:"-"`
	Recipe      FantasyExchangeRecipe `gorm:"foreignKey:RecipeID" json:"-"`
}
