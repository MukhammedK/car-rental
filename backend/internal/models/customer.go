package models

import "time"

type Customer struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	IIN       string    `json:"iin"`
	CreatedAt time.Time `json:"created_at"`
}
