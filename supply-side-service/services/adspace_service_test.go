package services

import (
	"auction-service/supply-side-service/database"
	"auction-service/supply-side-service/models"
	"auction-service/supply-side-service/repositories"
	"testing"
	"time"
)

func TestGetAllAdSpaces(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}

	var adSpaces []models.AdSpace
	repo := repositories.NewAdSpaceRepository(db)
	service := NewAdSpaceService(repo)

	adSpaces, err = service.GetAllAdSpaces()
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if len(adSpaces) <= 1 {
		t.Fatalf("no adspaces found")
	}
}

func TestCreateAdSpace(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}
	repo := repositories.NewAdSpaceRepository(db)
	service := NewAdSpaceService(repo)

	adSpace := models.AdSpace{
		Name:      "Test Ad Space",
		BasePrice: 100,
		EndTime:   time.Now().Add(time.Hour),
	}

	adSpaceID, err := service.CreateAdSpace(adSpace)
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if adSpaceID <= 0 {
		t.Fatalf("Invalid adSpaceID: %d", adSpaceID)
	}
}

func TestGetAdSpaceByID(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}

	adspaceID := 1
	repo := repositories.NewAdSpaceRepository(db)
	service := NewAdSpaceService(repo)

	adSpace, err := service.GetAdSpaceByID(adspaceID)
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if adSpace.Name == "" {
		t.Fatalf("adspace not found")
	}
}
func TestGetWinner(t *testing.T) {
	db, err := database.NewMySQLConnection()
	if err != nil {
		t.Fatal(err)
	}

	adspaceID := 1
	repo := repositories.NewAdSpaceRepository(db)
	service := NewAdSpaceService(repo)

	winnerID, err := service.GetWinner(adspaceID)
	if err != nil {
		t.Fatalf("CreateAdSpace failed: %v", err)
	}

	if winnerID <= 0 {
		t.Fatalf("Invalid adSpaceID: %d", winnerID)
	}
}
