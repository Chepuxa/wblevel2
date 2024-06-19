package domain

import "time"

// Event представляет собой структуру, хранящую информацию о событии
type Event struct {
	ID          int
	UserID      int
	Date        time.Time
	Description string
}
