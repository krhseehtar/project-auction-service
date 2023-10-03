package repositories

import (
	"auction-service/demand-side-service/models"
	"database/sql"
	"log"
	"time"
)

type BidRepository struct {
	db *sql.DB
}

func NewBidRepository(db *sql.DB) BidRepository {
	return BidRepository{db: db}
}

func (r *BidRepository) GetBidsByAdSpaceID(adSpaceID int) ([]models.Bid, error) {
	rows, err := r.db.Query("SELECT * FROM bids where ad_space_id=?", adSpaceID)
	if err != nil {
		log.Println("error in getBidsByAdSpaceID. error:", err)
		return nil, err
	}
	defer rows.Close()

	var bids []models.Bid
	var endTimeBytes []byte
	for rows.Next() {
		var bid models.Bid
		if err := rows.Scan(&bid.ID, &bid.AdSpaceID, &bid.BidderID, &bid.BidAmount, &endTimeBytes); err != nil {
			log.Println("error while processing result. error:", err)
			return nil, err
		}
		bid.Timestamp, err = time.Parse(time.DateTime, string(endTimeBytes))
		if err != nil {
			log.Println("error while converting endTimeBytes. error:", err)
			return nil, err
		}
		bids = append(bids, bid)
	}

	if err := rows.Err(); err != nil {
		log.Println("rows error. error:", err)
		return nil, err
	}

	return bids, nil
}

func (r *BidRepository) GetBidsByBidderID(bidderID int) ([]models.Bid, error) {
	rows, err := r.db.Query("SELECT * FROM bids where bidder_id=?", bidderID)
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return nil, err
	}
	defer rows.Close()

	var bids []models.Bid
	var endTimeBytes []byte
	for rows.Next() {
		var bid models.Bid
		if err := rows.Scan(&bid.ID, &bid.AdSpaceID, &bid.BidderID, &bid.BidAmount, &endTimeBytes); err != nil {
			log.Println("error while scanning query result query. error:", err)
			return nil, err
		}
		bid.Timestamp, err = time.Parse(time.DateTime, string(endTimeBytes))
		if err != nil {
			log.Println("error while converting endTime. error:", err)
			return nil, err
		}
		bids = append(bids, bid)
	}

	if err := rows.Err(); err != nil {
		log.Println("rows Error. error:", err)
		return nil, err
	}

	return bids, nil
}

func (r *BidRepository) GetAllBidsByBidderIDAndAdSpaceID(bidderID int, adspaceID int) ([]models.Bid, error) {
	rows, err := r.db.Query("SELECT * FROM bids where bidder_id=? and ad_space_id = ?", bidderID, adspaceID)
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return nil, err
	}
	defer rows.Close()

	var bids []models.Bid
	var endTimeBytes []byte
	for rows.Next() {
		var bid models.Bid
		if err := rows.Scan(&bid.ID, &bid.AdSpaceID, &bid.BidderID, &bid.BidAmount, &endTimeBytes); err != nil {
			log.Println("error while scanning query result query. error:", err)
			return nil, err
		}
		bid.Timestamp, err = time.Parse(time.DateTime, string(endTimeBytes))
		if err != nil {
			log.Println("error while converting endTime. error:", err)
			return nil, err
		}
		bids = append(bids, bid)
	}

	if err := rows.Err(); err != nil {
		log.Println("rows error. error:", err)
		return nil, err
	}

	return bids, nil
}

func (r *BidRepository) CreateBid(bid models.Bid) (int64, error) {
	result, err := r.db.Exec("INSERT INTO bids (ad_space_id, bidder_id, bid_amount, timestamp) VALUES (?, ?, ?, ?)",
		bid.AdSpaceID, bid.BidderID, bid.BidAmount, bid.Timestamp)
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return -1, err
	}
	var bidID int64
	bidID, err = result.LastInsertId()
	if err != nil {
		log.Println("error while fetching lastInsertId. error:", err)
		return -1, err
	}

	return bidID, nil
}

func (r *BidRepository) CreateBidder(bidder models.Bidder) (int64, error) {
	result, err := r.db.Exec("INSERT INTO bidders (name, email) VALUES (?, ?)",
		bidder.Name, bidder.Email)
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return -1, err
	}
	var bidderID int64
	bidderID, err = result.LastInsertId()
	if err != nil {
		log.Println("error while fetching lastInsertId. error:", err)
		return -1, err
	}

	return bidderID, nil
}

func (r *BidRepository) GetBidderById(bidderID int) (models.Bidder, error) {
	var bidder models.Bidder
	err := r.db.QueryRow("SELECT id, name, email from bidders WHERE id = ?", bidderID).
		Scan(&bidder.ID, &bidder.Name, &bidder.Email)

	if err != nil {
		log.Println("error while excuting query. error:", err)
		return models.Bidder{}, err
	}

	return bidder, nil
}

func (r *BidRepository) GetAllBidders() ([]models.Bidder, error) {
	rows, err := r.db.Query("SELECT * FROM bidders")
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return nil, err
	}
	defer rows.Close()

	var bidders []models.Bidder
	for rows.Next() {
		var bidder models.Bidder
		if err := rows.Scan(&bidder.ID, &bidder.Name, &bidder.Email); err != nil {
			log.Println("error while scanning query result query. error:", err)
			return nil, err
		}
		bidders = append(bidders, bidder)
	}

	if err := rows.Err(); err != nil {
		log.Println("rows error. error:", err)
		return nil, err
	}

	return bidders, nil
}

func (r *BidRepository) GetBidderByEmailId(emailID string) (int64, error) {
	var rowsAffected int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM bidders WHERE email = ?", emailID).Scan(&rowsAffected)

	if err != nil {
		log.Println("error while excuting query. error:", err)
		return -1, err
	}

	return rowsAffected, nil
}

func (r *BidRepository) AdSpaceExists(adSpaceID int) (bool, error) {
	var rowsAffected int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM ad_spaces WHERE id = ?", adSpaceID).Scan(&rowsAffected)

	if err != nil {
		log.Println("error while excuting query. error:", err)
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *BidRepository) BidderExists(BidderID int) (bool, error) {
	var rowsAffected int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM bidders WHERE id = ?", BidderID).Scan(&rowsAffected)
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *BidRepository) IsActive(adSpaceID int) (bool, error) {
	var endTime time.Time
	var endTimeBytes []byte
	err := r.db.QueryRow("SELECT end_time FROM ad_spaces WHERE id = ?", adSpaceID).Scan(&endTimeBytes)
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return false, err
	}
	endTime, err = time.Parse(time.DateTime, string(endTimeBytes))
	if err != nil {
		log.Println("error while converting endTime. error:", err)
		return false, err
	}
	currentTimestamp := time.Now().UTC()
	if currentTimestamp.After(endTime) {
		return false, nil
	}

	return true, nil
}

func (r *BidRepository) IsValidBidAmount(bid models.Bid) (bool, error) {
	var basePrice float64
	var currentBid float64
	err := r.db.QueryRow("SELECT base_price, current_bid FROM ad_spaces WHERE id = ?", bid.AdSpaceID).Scan(&basePrice, &currentBid)
	if err != nil {
		log.Println("error while excuting query. error:", err)
		return false, err
	}

	if bid.BidAmount >= basePrice && bid.BidAmount > currentBid {
		return true, nil
	}

	return false, nil
}

func (r *BidRepository) UpdateCurrentBid(bid models.Bid) (bool, error) {
	stmt, err := r.db.Prepare("UPDATE ad_spaces SET current_bid = ? WHERE id = ?")
	if err != nil {
		log.Println("error while preparing update query. error:", err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bid.BidAmount, bid.AdSpaceID)
	if err != nil {
		log.Println("error while executing update query. error:", err)
		return false, err
	}

	return true, nil
}
