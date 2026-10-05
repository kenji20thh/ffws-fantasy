package db

import (
	"log"

	"ffws/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// defaultRarities are the initial card tiers, lowest to highest.
var defaultRarities = []models.FantasyRarity{
	{Name: "Common", Slug: "common", Color: "#9CA3AF", SortOrder: 1},
	{Name: "Uncommon", Slug: "uncommon", Color: "#22C55E", SortOrder: 2},
	{Name: "Rare", Slug: "rare", Color: "#3B82F6", SortOrder: 3},
	{Name: "Epic", Slug: "epic", Color: "#A855F7", SortOrder: 4},
	{Name: "Legendary", Slug: "legendary", Color: "#F59E0B", SortOrder: 5},
	{Name: "Mythic", Slug: "mythic", Color: "#EF4444", SortOrder: 6},
}

// SeedFantasyRarities inserts the initial rarities. It is safe to run on every startup:
// rows are matched by slug and existing ones are left untouched, so later edits to a
// rarity's name, colour or order are never overwritten.
func SeedFantasyRarities(database *gorm.DB) error {
	for _, r := range defaultRarities {
		rarity := r
		if err := database.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "slug"}},
			DoNothing: true,
		}).Create(&rarity).Error; err != nil {
			return err
		}
	}
	log.Println("fantasy rarities seeded")
	return nil
}
