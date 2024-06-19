package domain

import (
	"context"
	"time"
)

// Repository представляет прослойку CRUD-операций
type Repository interface {
	CreateEvent(ctx context.Context, event Event) (int, error)
	UpdateEvent(ctx context.Context, event Event) error
	DeleteEvent(ctx context.Context, eventID int) error
	GetEventByID(ctx context.Context, eventID int) (Event, error)
	GetEventsForDay(ctx context.Context, date time.Time) ([]Event, error)
	GetEventsForWeek(ctx context.Context, date time.Time) ([]Event, error)
	GetEventsForMonth(ctx context.Context, date time.Time) ([]Event, error)
}
