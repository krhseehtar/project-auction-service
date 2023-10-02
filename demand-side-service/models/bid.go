package models

import "time"

type Bid struct {
	ID        int       `json:"id"`
	AdSpaceID int       `json:"adSpaceID"`
	BidderID  int       `json:"bidderID"`
	BidAmount float64   `json:"bidAmount"`
	Timestamp time.Time `json:"timestamp"`
}
