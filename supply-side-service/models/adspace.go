package models

import "time"

type AdSpace struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	BasePrice  float64   `json:"basePrice"`
	EndTime    time.Time `json:"endTime"`
	CurrentBid float64   `json:"currentBid"`
	WinnerID   int       `json:"winnerID"`
}
