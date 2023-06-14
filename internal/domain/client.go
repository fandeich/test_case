package domain

import "time"

type Client struct {
	ID        string
	Name      string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
