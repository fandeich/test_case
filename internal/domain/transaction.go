package domain

import "time"

type Transaction struct {
	ID        string
	ClientID  string
	Amount    float64
	Status    TransactionStatus
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionStatus int

const (
	Pending TransactionStatus = iota
	Completed
	Rejected
)

func (status TransactionStatus) String() string {
	switch status {
	case Pending:
		return "Pending"
	case Completed:
		return "Completed"
	case Rejected:
		return "Rejected"
	default:
		return "Unknown"
	}
}
