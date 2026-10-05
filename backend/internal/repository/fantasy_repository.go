package repository

import (
	"errors"

	"ffws/internal/models"

	"gorm.io/gorm"
)

// FantasyRepository handles fantasy teams (the user's fantasy club for a tournament).
//
// The old direct-selection system (pick any 4 players from the global pool, chips, budget)
// has been removed; the card/pack system replaces it in later phases.
type FantasyRepository struct {
	db *gorm.DB
}

func NewFantasyRepository(db *gorm.DB) *FantasyRepository {
	return &FantasyRepository{db: db}
}

var (
	ErrFantasyTeamExists   = errors.New("you already have a fantasy team for this tournament")
	ErrFantasyTeamNotFound = errors.New("fantasy team not found")

	// ErrDayNotFound is shared with the Prediction repository/handler.
	ErrDayNotFound = errors.New("tournament day not found")
)

func (r *FantasyRepository) CreateTeam(userID, tournamentID uint, teamName, country string) (*models.FantasyTeam, error) {
	var count int64
	r.db.Model(&models.FantasyTeam{}).Where("user_id = ? AND tournament_id = ?", userID, tournamentID).Count(&count)
	if count > 0 {
		return nil, ErrFantasyTeamExists
	}
	team := models.FantasyTeam{UserID: userID, TournamentID: tournamentID, TeamName: teamName, Country: country}
	if err := r.db.Create(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *FantasyRepository) GetTeamByUser(userID, tournamentID uint) (*models.FantasyTeam, error) {
	var team models.FantasyTeam
	err := r.db.Where("user_id = ? AND tournament_id = ?", userID, tournamentID).First(&team).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrFantasyTeamNotFound
		}
		return nil, err
	}
	return &team, nil
}

func (r *FantasyRepository) GetTeamByID(fantasyTeamID uint) (*models.FantasyTeam, error) {
	var team models.FantasyTeam
	err := r.db.First(&team, fantasyTeamID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrFantasyTeamNotFound
		}
		return nil, err
	}
	return &team, nil
}

type FantasyStanding struct {
	FantasyTeamID uint   `json:"fantasy_team_id"`
	TeamName      string `json:"team_name"`
	Country       string `json:"country"`
	Points        int    `json:"points"`
}

// GetStandings lists every fantasy team of the tournament.
//
// Points are 0 for now: they used to come from the old direct-selection tables, which no longer
// feed the game. Card-based scoring arrives in a later phase. dayID is accepted so the API
// contract stays the same, but has no effect until then.
func (r *FantasyRepository) GetStandings(tournamentID uint, dayID *uint) ([]FantasyStanding, error) {
	var teams []models.FantasyTeam
	if err := r.db.Where("tournament_id = ?", tournamentID).Order("id ASC").Find(&teams).Error; err != nil {
		return nil, err
	}

	standings := make([]FantasyStanding, 0, len(teams))
	for _, t := range teams {
		standings = append(standings, FantasyStanding{
			FantasyTeamID: t.ID, TeamName: t.TeamName, Country: t.Country, Points: 0,
		})
	}
	return standings, nil
}
