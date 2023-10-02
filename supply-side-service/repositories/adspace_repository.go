package repositories

import (
	"auction-service/supply-side-service/models"
	"database/sql"
	"errors"
	"time"
)

type AdSpaceRepository struct {
	db *sql.DB
}

func NewAdSpaceRepository(db *sql.DB) AdSpaceRepository {
	return AdSpaceRepository{db: db}
}

func (r *AdSpaceRepository) GetAllAdSpaces() ([]models.AdSpace, error) {
	rows, err := r.db.Query("SELECT * FROM ad_spaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var adSpaces []models.AdSpace
	for rows.Next() {
		var adSpace models.AdSpace
		var endTimeBytes []byte
		if err := rows.Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice, &endTimeBytes, &adSpace.CurrentBid, &adSpace.WinnerID); err != nil {
			return nil, err
		}
		adSpace.EndTime, err = time.Parse("2006-01-02 15:04:05", string(endTimeBytes))
		adSpaces = append(adSpaces, adSpace)
		if err != nil {
			// Handle error (e.g., ad space not found)
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return adSpaces, nil
}

func (r *AdSpaceRepository) GetAdSpaceByID(id int) (models.AdSpace, error) {
	var adSpace models.AdSpace
	var endTimeBytes []byte
	err := r.db.QueryRow("SELECT id, name, base_price, end_time, current_bid, winner_id FROM ad_spaces WHERE id = ?", id).
		Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice, &endTimeBytes, &adSpace.CurrentBid, &adSpace.WinnerID)

	if err != nil {
		// Handle error (e.g., ad space not found)
		return models.AdSpace{}, err
	}
	adSpace.EndTime, err = time.Parse("2006-01-02 15:04:05", string(endTimeBytes))
	if err != nil {
		// Handle error (e.g., ad space not found)
		return models.AdSpace{}, err
	}

	return adSpace, nil
}

func (r *AdSpaceRepository) CreateAdSpace(adSpace models.AdSpace) (int64, error) {
	result, err := r.db.Exec("INSERT INTO ad_spaces (name, base_price, end_time, current_bid, winner_id) VALUES (?, ?, ?, ?, ?)",
		adSpace.Name, adSpace.BasePrice, adSpace.EndTime, adSpace.CurrentBid, adSpace.WinnerID)
	if err != nil {
		return -1, err
	}

	var adSpaceId int64
	adSpaceId, err = result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return adSpaceId, nil

}

func (r *AdSpaceRepository) GetWinner(id int) (int, error) {
	var adSpace models.AdSpace
	var endTimeBytes []byte
	err := r.db.QueryRow("SELECT id, name, base_price, end_time, current_bid, winner_id FROM ad_spaces WHERE id = ?", id).
		Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice, &endTimeBytes, &adSpace.CurrentBid, &adSpace.WinnerID)

	if err != nil {
		// Handle error (e.g., ad space not found)
		return -1, err
	}
	adSpace.EndTime, err = time.Parse("2006-01-02 15:04:05", string(endTimeBytes))
	if err != nil {
		return -1, err
	}
	if adSpace.EndTime.After(time.Now().UTC()) {
		return -1, errors.New("auction in progress")
	}

	return adSpace.WinnerID, nil
}

func (r *AdSpaceRepository) FindWinner(adspaceID int) (int, error) {
	var winnerID int
	err := r.db.QueryRow("SELECT bidder_id FROM bids WHERE bid_amount = ( SELECT MAX(bid_amount) FROM bids WHERE ad_space_id = ? )AND ad_space_id = ?", adspaceID, adspaceID).
		Scan(&winnerID)

	if err != nil {
		// Handle error (e.g., ad space not found)
		return -1, err
	}

	return winnerID, nil
}

func (r *AdSpaceRepository) UpdateWinner(adspaceID int, winnerID int) (bool, error) {
	stmt, err := r.db.Prepare("UPDATE ad_spaces SET winner = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(winnerID, adspaceID)
	if err != nil {
		return false, err
	}

	return true, nil
}
