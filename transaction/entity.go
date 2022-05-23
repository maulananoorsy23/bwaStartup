package transaction

import (
	"time"
)

type Transaction struct {
	ID         int
	campaignID int
	UserID     int
	Amount     int
	Status     int
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
