package services

import (
	"auction-service/supply-side-service/models"
	"auction-service/supply-side-service/repositories"
	"errors"
	"log"
	"time"
)

type AdSpaceService interface {
	GetAllAdSpaces() ([]models.AdSpace, error)
	GetAdSpaceByID(id int) (models.AdSpace, error)
	CreateAdSpace(adSpace models.AdSpace) (int64, error)
	GetWinner(id int) (int, error)
}

type adSpaceService struct {
	repo repositories.AdSpaceRepository
}

// NewAdSpaceService creates a new AdSpaceService instance with the provided repository.
func NewAdSpaceService(repo repositories.AdSpaceRepository) AdSpaceService {
	return &adSpaceService{repo: repo}
}

// GetAllAdSpaces retrieves all ad spaces from the repository.
func (s *adSpaceService) GetAllAdSpaces() ([]models.AdSpace, error) {
	return s.repo.GetAllAdSpaces()
}

// GetAdSpaceByID retrieves an ad space by its ID from the repository.
func (s *adSpaceService) GetAdSpaceByID(id int) (models.AdSpace, error) {
	return s.repo.GetAdSpaceByID(id)
}

// CreateAdSpace creates a new ad space in the repository with the provided details.
func (s *adSpaceService) CreateAdSpace(adSpace models.AdSpace) (int64, error) {
	currentTimestamp := time.Now().UTC()

	if currentTimestamp.After(adSpace.EndTime) {
		return -1, errors.New("the end time of the auction cannot be less than the current time")
	}
	if adSpace.BasePrice < 0 {
		return -1, errors.New("invalid base price")
	}
	if len(adSpace.Name) == 0 {
		return -1, errors.New("name cannot be empty")
	}
	return s.repo.CreateAdSpace(adSpace)
}

// GetWinner retrieves the winner of the specified ad space auction, handling both direct winner retrieval
// and finding the winner if not directly available in the database.
func (s *adSpaceService) GetWinner(adSpaceID int) (int, error) {
	var winnerID int
	var err error

	winnerID, err = s.repo.GetWinner(adSpaceID)
	if err != nil {
		log.Println("error in getWinner(). error:", err)
		return -1, err
	}

	if winnerID == 0 {
		winnerID, err = s.repo.FindWinner(adSpaceID)
	}
	if err != nil {
		log.Println("error in findWinner(). error:", err)
		return -1, err
	}

	_, err = s.repo.UpdateWinner(adSpaceID, winnerID)
	if err != nil {
		log.Println("error in updateWinner(). error:", err)
		return -1, err
	}
	return winnerID, nil

}
