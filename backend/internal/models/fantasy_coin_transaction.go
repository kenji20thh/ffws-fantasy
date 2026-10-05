package models

import "time"

// Coin transaction types. Only the vocabulary is defined here; none of these flows exist yet.
const (
	CoinTxStarter       = "starter"
	CoinTxDailyLogin    = "daily_login"
	CoinTxMatchReward   = "match_reward"
	CoinTxDuplicateSale = "duplicate_sale"
	CoinTxPackPurchase  = "pack_purchase"
	CoinTxExchange      = "exchange"
	CoinTxChallenge     = "challenge"
	CoinTxLeagueReward  = "league_reward"
)

// FantasyCoinTransaction is the audit log of every coin change. Amount is signed
// (positive = coins earned, negative = coins spent).
type FantasyCoinTransaction struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	FantasyTeamID uint   `gorm:"not null;index:idx_coin_tx_team_created,priority:1" json:"fantasy_team_id"`
	Amount        int    `gorm:"not null" json:"amount"`
	Type          string `gorm:"not null;index" json:"type"`

	// ReferenceID points at whatever caused the change (a pack opening, an exchange, a match ...),
	// interpreted according to Type. It is NULL for types that have no source row, such as "starter".
	ReferenceID *uint `json:"reference_id"`

	CreatedAt time.Time `gorm:"index:idx_coin_tx_team_created,priority:2" json:"created_at"`

	FantasyTeam FantasyTeam `gorm:"foreignKey:FantasyTeamID" json:"-"`
}
