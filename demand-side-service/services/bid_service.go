package services

import (
	"auction-service/demand-side-service/models"
	"auction-service/demand-side-service/repositories"
	"errors"
	"log"
	"time"
)

type BidService interface {
	CreateBidder(bidder models.Bidder) (int64, error)
	GetAllBidders() ([]models.Bidder, error)
	GetBidderById(bidderID int) (models.Bidder, error)
	PlaceBid(bid models.Bid) (int64, error)
	GetBidsByAdSpaceID(adSpaceID int) ([]models.Bid, error)
	GetAllBidsByBidderID(bidderID int) ([]models.Bid, error)
	GetAllBidsByBidderIDAndAdSpaceID(bidderID int, adspaceID int) ([]models.Bid, error)
}

type bidService struct {
	repo repositories.BidRepository
}

// NewBidService creates a new BidService instance with the provided repository.
func NewBidService(repo repositories.BidRepository) BidService {
	return &bidService{repo: repo}
}

// CreateBidder creates a new bidder and validates the input data.
func (s *bidService) CreateBidder(bidder models.Bidder) (int64, error) {

	if len(bidder.Email) == 0 {
		return -1, errors.New("invalid email id")
	}
	if len(bidder.Name) == 0 {
		return -1, errors.New("name cannot be empty")
	}

	rowCount, err := s.repo.GetBidderByEmailId(bidder.Email)
	if err != nil {
		log.Println("error in getBidderByEmailId(). error:", err)
		return -1, err
	}

	if rowCount > 0 {
		return -1, errors.New("email id already exists")
	}

	return s.repo.CreateBidder(bidder)
}

// GetBidderById retrieves a bidder by their ID.
func (s *bidService) GetBidderById(bidderID int) (models.Bidder, error) {
	return s.repo.GetBidderById(bidderID)
}

// GetAllBidders retrieves all registered bidders.
func (s *bidService) GetAllBidders() ([]models.Bidder, error) {
	return s.repo.GetAllBidders()
}

// PlaceBid places a new bid for a specific ad space and validates the bid data.
func (s *bidService) PlaceBid(bid models.Bid) (int64, error) {
	var bidderExists, AdSpaceExists, isActive, isValidBidAmount bool
	var err error
	bidderExists, err = s.repo.BidderExists(bid.BidderID)
	if err != nil {
		log.Println("error in bidderExists(). error:", err)
		return -1, err
	}
	if !bidderExists {
		return -1, errors.New("invalid bidderID")
	}

	AdSpaceExists, err = s.repo.AdSpaceExists(bid.AdSpaceID)
	if err != nil {
		log.Println("error in adSpaceExists(). error:", err)
		return -1, err
	}
	if !AdSpaceExists {
		return -1, errors.New("invalid AdSpaceID")
	}

	isActive, err = s.repo.IsActive(bid.AdSpaceID)
	if err != nil {
		log.Println("error in isActive(). error:", err)
		return -1, err
	}
	if !isActive {
		return -1, errors.New("auction not active")
	}

	isValidBidAmount, err = s.repo.IsValidBidAmount(bid)
	if err != nil {
		log.Println("error in isValidBidAmount(). error:", err)
		return -1, err
	}
	if !isValidBidAmount {
		return -1, errors.New("bid amount is less than current/base price")
	}

	_, err = s.repo.UpdateCurrentBid(bid)
	if err != nil {
		log.Println("error in updateCurrentBid(). error:", err)
		return -1, err
	}

	bid.Timestamp = time.Now().UTC()
	return s.repo.CreateBid(bid)
}

// GetBidsByAdSpaceID retrieves bids for a specific ad space using its ID.
func (s *bidService) GetBidsByAdSpaceID(adSpaceID int) ([]models.Bid, error) {
	return s.repo.GetBidsByAdSpaceID(adSpaceID)
}

// GetAllBidsByBidderID retrieves bids placed by a specific bidder using their ID.
func (s *bidService) GetAllBidsByBidderID(bidderID int) ([]models.Bid, error) {
	return s.repo.GetBidsByBidderID(bidderID)
}

// GetAllBidsByBidderIDAndAdSpaceID retrieves bids placed by a specific bidder for a specific ad space using their IDs.
func (s *bidService) GetAllBidsByBidderIDAndAdSpaceID(bidderID int, adspaceID int) ([]models.Bid, error) {
	return s.repo.GetAllBidsByBidderIDAndAdSpaceID(bidderID, adspaceID)
}
