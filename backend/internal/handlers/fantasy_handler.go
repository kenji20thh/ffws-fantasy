package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"ffws/internal/repository"

	"github.com/gin-gonic/gin"
)

type FantasyHandler struct {
	repo *repository.FantasyRepository
}

func NewFantasyHandler(repo *repository.FantasyRepository) *FantasyHandler {
	return &FantasyHandler{repo: repo}
}

type createFantasyTeamRequest struct {
	TournamentID uint   `json:"tournament_id" binding:"required"`
	TeamName     string `json:"team_name" binding:"required"`
	Country      string `json:"country"`
}

func (h *FantasyHandler) CreateTeam(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req createFantasyTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.repo.CreateTeam(userID, req.TournamentID, req.TeamName, req.Country)
	if err != nil {
		if errors.Is(err, repository.ErrFantasyTeamExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create fantasy team"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": team})
}

func (h *FantasyHandler) GetMyTeam(c *gin.Context) {
	userID := c.GetUint("user_id")
	tournamentID, err := strconv.ParseUint(c.Query("tournament_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tournament_id query param is required"})
		return
	}

	team, err := h.repo.GetTeamByUser(userID, uint(tournamentID))
	if err != nil {
		if errors.Is(err, repository.ErrFantasyTeamNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch fantasy team"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": team})
}

func (h *FantasyHandler) GetStandings(c *gin.Context) {
	tournamentID, err := strconv.ParseUint(c.Query("tournament_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tournament_id query param is required"})
		return
	}

	var dayIDPtr *uint
	if dayIDStr := c.Query("day_id"); dayIDStr != "" {
		dayID, err := strconv.ParseUint(dayIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid day_id"})
			return
		}
		d := uint(dayID)
		dayIDPtr = &d
	}

	standings, err := h.repo.GetStandings(uint(tournamentID), dayIDPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch standings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": standings})
}

// GetTeamProfile is the public view of a fantasy team. It exposes the team's identity only;
// squads and collections will be added by later phases. The owner's user id is never exposed.
func (h *FantasyHandler) GetTeamProfile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fantasy team id"})
		return
	}

	team, err := h.repo.GetTeamByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "fantasy team not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"team": gin.H{
		"id":            team.ID,
		"tournament_id": team.TournamentID,
		"team_name":     team.TeamName,
		"country":       team.Country,
		"created_at":    team.CreatedAt,
	}}})
}
