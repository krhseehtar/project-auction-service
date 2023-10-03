package services

import (
	"auction-service/demand-side-service/database"
	"auction-service/demand-side-service/models"
	"auction-service/demand-side-service/repositories"
	"testing"
)

func TestCreateBidder(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}

	bidder := models.Bidder{
		Name:  "Test Bidder",
		Email: "example1@gmail.com",
	}
	repo := repositories.NewBidRepository(db)
	service := NewBidService(repo)

	bidderId, err := service.CreateBidder(bidder)
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if bidderId <= 0 {
		t.Fatalf("Invalid adSpaceID: %d", bidderId)
	}
}

func TestGetAllBidders(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}

	var adSpaces []models.Bidder
	repo := repositories.NewBidRepository(db)
	service := NewBidService(repo)

	adSpaces, err = service.GetAllBidders()
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if len(adSpaces) <= 1 {
		t.Fatalf("no adspaces found")
	}
}

func TestPlaceBid(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}

	bid := models.Bid{
		AdSpaceID: 1,
		BidderID:  1,
		BidAmount: 110,
	}
	repo := repositories.NewBidRepository(db)
	service := NewBidService(repo)

	bidId, err := service.PlaceBid(bid)
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if bidId <= 0 {
		t.Fatalf("Invalid adSpaceID: %d", bidId)
	}
}

func TestGetBidsByAdSpaceID(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}

	var bids []models.Bid
	adSpaceID := 1
	repo := repositories.NewBidRepository(db)
	service := NewBidService(repo)

	bids, err = service.GetBidsByAdSpaceID(adSpaceID)
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if len(bids) <= 1 {
		t.Fatalf("no adspaces found")
	}
}
