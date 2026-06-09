package models

import "time"

type Booking struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	CarID      int       `json:"car_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	Source     string    `json:"source"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
