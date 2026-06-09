package models

import "time"

type Car struct {
	ID           int       `json:"id"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	LicensePlate string    `json:"license_plate"`
	Color        string    `json:"color"`
	Transmission string    `json:"transmission"`
	FuelType     string    `json:"fuel_type"`
	DailyPrice   float64   `json:"daily_price"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}
