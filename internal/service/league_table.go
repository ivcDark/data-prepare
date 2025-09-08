package service

import (
	"log"
	"data-preparer/internal/repository"
)

type LeagueTableService struct {
	repo *repository.MySQLRepository
}

func NewLeagueTableService(repo *repository.MySQLRepository) *LeagueTableService {
	return &LeagueTableService{repo: repo}
}

func (s *LeagueTableService) UpdateAllTables() {
	leagues, err := s.repo.GetAllLeagueSeasons()
	if err != nil {
		log.Printf("Error getting leagues: %v", err)
		return
	}

	for _, league := range leagues {
		log.Printf("Updating table for league: %s", league.Name)
		// Логика обновления таблицы
		// ...
	}
}