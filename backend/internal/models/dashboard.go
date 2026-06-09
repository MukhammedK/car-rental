package models

type DashboardStats struct {
	CarsTotal      int     `json:"cars_total"`
	CarsAvailable  int     `json:"cars_available"`
	BookingsNew    int     `json:"bookings_new"`
	BookingsActive int     `json:"bookings_active"`
	Revenue        float64 `json:"revenue"`
}
