package services

import (
	"auction-service/supply-side-service/models"
	"auction-service/supply-side-service/repositories"
	"errors"
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

func NewAdSpaceService(repo repositories.AdSpaceRepository) AdSpaceService {
	return &adSpaceService{repo: repo}
}

func (s *adSpaceService) GetAllAdSpaces() ([]models.AdSpace, error) {
	return s.repo.GetAllAdSpaces()
}

func (s *adSpaceService) GetAdSpaceByID(id int) (models.AdSpace, error) {
	return s.repo.GetAdSpaceByID(id)
}

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
func (s *adSpaceService) GetWinner(adSpaceID int) (int, error) {
	var winnerID int
	var err error

	winnerID, err = s.repo.GetWinner(adSpaceID)
	if err != nil {
		return -1, err
	}

	if winnerID == 0 {
		winnerID, err = s.repo.FindWinner(adSpaceID)
	}
	if err != nil {
		return -1, err
	}

	_, err = s.repo.UpdateWinner(adSpaceID, winnerID)
	if err != nil {
		return -1, err
	}
	return winnerID, nil

}
