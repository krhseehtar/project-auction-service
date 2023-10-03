package repositories

import (
	"auction-service/supply-side-service/models"
	"database/sql"
	"errors"
	"log"
	"time"
)

type AdSpaceRepository struct {
	db *sql.DB
}

// NewAdSpaceRepository creates a new AdSpaceRepository instance with the provided database connection.
func NewAdSpaceRepository(db *sql.DB) AdSpaceRepository {
	return AdSpaceRepository{db: db}
}

// GetAllAdSpaces retrieves all ad spaces from the database.
func (r *AdSpaceRepository) GetAllAdSpaces() ([]models.AdSpace, error) {
	rows, err := r.db.Query("SELECT * FROM ad_spaces")
	if err != nil {
		log.Println("error while executing query. error:", err)
		return nil, err
	}
	defer rows.Close()

	var adSpaces []models.AdSpace
	for rows.Next() {
		var adSpace models.AdSpace
		var endTimeBytes []byte
		if err := rows.Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice, &endTimeBytes, &adSpace.CurrentBid, &adSpace.WinnerID); err != nil {
			log.Println("error while scanning query result. error:", err)
			return nil, err
		}
		adSpace.EndTime, err = time.Parse(time.DateTime, string(endTimeBytes))
		adSpaces = append(adSpaces, adSpace)
		if err != nil {
			log.Println("error while converting endTimeBytes. error:", err)
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Println("rows.err(). error:", err)
		return nil, err
	}

	return adSpaces, nil
}

// GetAdSpaceByID retrieves an ad space by its ID from the database.
func (r *AdSpaceRepository) GetAdSpaceByID(id int) (models.AdSpace, error) {
	var adSpace models.AdSpace
	var endTimeBytes []byte
	err := r.db.QueryRow("SELECT id, name, base_price, end_time, current_bid, winner_id FROM ad_spaces WHERE id = ?", id).
		Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice, &endTimeBytes, &adSpace.CurrentBid, &adSpace.WinnerID)

	if err != nil {
		log.Println("error while executing query. error:", err)
		return models.AdSpace{}, err
	}
	adSpace.EndTime, err = time.Parse(time.DateTime, string(endTimeBytes))
	if err != nil {
		log.Println("error while converting endTimeBytes. error:", err)
		return models.AdSpace{}, err
	}

	return adSpace, nil
}

// CreateAdSpace creates a new ad space in the database and returns its ID.
func (r *AdSpaceRepository) CreateAdSpace(adSpace models.AdSpace) (int64, error) {
	result, err := r.db.Exec("INSERT INTO ad_spaces (name, base_price, end_time, current_bid, winner_id) VALUES (?, ?, ?, ?, ?)",
		adSpace.Name, adSpace.BasePrice, adSpace.EndTime, adSpace.CurrentBid, adSpace.WinnerID)
	if err != nil {
		log.Println("error while executing query. error:", err)
		return -1, err
	}

	adSpaceId, err := result.LastInsertId()
	if err != nil {
		log.Println("error while fetching lastInsertId. error:", err)
		return -1, err
	}

	return adSpaceId, nil

}

// GetWinner retrieves the winner of an ad space auction based on its ID.
func (r *AdSpaceRepository) GetWinner(id int) (int, error) {
	var adSpace models.AdSpace
	var endTimeBytes []byte
	err := r.db.QueryRow("SELECT id, name, base_price, end_time, current_bid, winner_id FROM ad_spaces WHERE id = ?", id).
		Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice, &endTimeBytes, &adSpace.CurrentBid, &adSpace.WinnerID)

	if err != nil {
		log.Println("error while executing query. error:", err)
		return -1, errors.New("ad-space not found")
	}
	adSpace.EndTime, err = time.Parse(time.DateTime, string(endTimeBytes))
	if err != nil {
		log.Println("error while converting endTimeBytes. error:", err)
		return -1, err
	}
	if adSpace.EndTime.After(time.Now().UTC()) {
		return -1, errors.New("auction in progress")
	}

	return adSpace.WinnerID, nil
}

// FindWinner finds the winner of an ad space auction based on the highest bid amount.
func (r *AdSpaceRepository) FindWinner(adspaceID int) (int, error) {
	var winnerID int
	err := r.db.QueryRow("SELECT bidder_id FROM bids WHERE bid_amount = ( SELECT MAX(bid_amount) FROM bids WHERE ad_space_id = ? )AND ad_space_id = ?", adspaceID, adspaceID).
		Scan(&winnerID)

	if err != nil {
		log.Println("error while executing query. error:", err)
		return -1, errors.New("no bids found for this ad-space")
	}

	return winnerID, nil
}

// UpdateWinner updates the winner of an ad space in the database.
func (r *AdSpaceRepository) UpdateWinner(adspaceID int, winnerID int) (bool, error) {
	stmt, err := r.db.Prepare("UPDATE ad_spaces SET winner_id = ? WHERE id = ?")
	if err != nil {
		log.Println("error while preparing update query. error:", err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(winnerID, adspaceID)
	if err != nil {
		log.Println("error while executing update query. error:", err)
		return false, err
	}

	return true, nil
}
