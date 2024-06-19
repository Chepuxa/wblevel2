package adapters

import (
	"context"
	"fmt"
	"sync"
	"time"
	"webtech/http/internal/calendar/domain"
)

// СacheEventRepository представляет собой кеш событий
type СacheEventRepository struct {
	cache         map[int]domain.Event
	autoIncrement int
	maxSize       int
	mu            *sync.RWMutex
}

// NewCacheEventRepository отвечает за иницализвацию cacheEventRepository
func NewCacheEventRepository(maxSize int) *СacheEventRepository {
	return &СacheEventRepository{
		cache:         make(map[int]domain.Event, maxSize),
		autoIncrement: 1,
		maxSize:       maxSize,
		mu:            &sync.RWMutex{},
	}
}

// CreateEvent отвечает за добавление события в кеш
func (r *СacheEventRepository) CreateEvent(ctx context.Context, domainEvent domain.Event) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.autoIncrement == r.maxSize {
		r.autoIncrement = 1
	}

	if len(r.cache) == r.maxSize {
		delete(r.cache, r.autoIncrement)
	}

	if _, ok := r.cache[domainEvent.ID]; !ok {
		domainEvent.ID = r.autoIncrement
		r.autoIncrement++
	}

	r.cache[domainEvent.ID] = domainEvent

	return domainEvent.ID, nil
}

// GetEventByID отвечает за получение события из кеша
func (r *СacheEventRepository) GetEventByID(ctx context.Context, eventID int) (domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	event, ok := r.cache[eventID]
	if ok {
		return event, nil
	}
	return domain.Event{}, fmt.Errorf("error: can't find event")
}

// UpdateEvent отвечает за обновление события из кеша
func (r *СacheEventRepository) UpdateEvent(ctx context.Context, updatedEvent domain.Event) error {
	_, err := r.GetEventByID(ctx, updatedEvent.ID)
	if err != nil {
		return err
	}

	_, err = r.CreateEvent(ctx, updatedEvent)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEvent отвечает за удаление события из кеша
func (r *СacheEventRepository) DeleteEvent(ctx context.Context, eventID int) error {
	_, err := r.GetEventByID(ctx, eventID)
	if err != nil {
		return err
	}

	r.mu.Lock()
	delete(r.cache, eventID)
	r.mu.Unlock()

	return nil
}

// GetEventsForDay отвечает за получение события по дню из кеша
func (r *СacheEventRepository) GetEventsForDay(ctx context.Context, date time.Time) ([]domain.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	eventsForDay := make([]domain.Event, 0, 5)

	for _, v := range r.cache {
		if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() && v.Date.Day() == date.Day() {
			eventsForDay = append(eventsForDay, v)
		}
	}

	return eventsForDay, nil
}

// GetEventsForWeek отвечает за получение событий по неделе из кеша
func (r *СacheEventRepository) GetEventsForWeek(ctx context.Context, date time.Time) ([]domain.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	eventsForWeek := make([]domain.Event, 0, 10)

	weekday := int(date.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	startOfWeek := date.AddDate(0, 0, -weekday)
	endOfWeek := startOfWeek.AddDate(0, 0, 8)

	for _, v := range r.cache {
		if v.Date.After(startOfWeek) && v.Date.Before(endOfWeek) {
			eventsForWeek = append(eventsForWeek, v)
		}
	}

	return eventsForWeek, nil
}

// GetEventsForMonth отвечает за получение событий по месяцу из кеша
func (r *СacheEventRepository) GetEventsForMonth(ctx context.Context, date time.Time) ([]domain.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	eventsForMonth := make([]domain.Event, 0, 20)

	for _, v := range r.cache {
		if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
			eventsForMonth = append(eventsForMonth, v)
		}
	}

	return eventsForMonth, nil
}
