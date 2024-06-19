package usecase

import (
	"context"
	"time"
	"webtech/http/internal/calendar/domain"
)

// GetEventsForMonthUseCase представляет обработчик бизнес-логики для операции получения события по месяцу
type GetEventsForMonthUseCase struct {
	eventRepository domain.Repository
}

// NewGetEventsForMonthUseCase отвечает за Инициализацию обработчика бизнес-логики для операции получения события по месяцу
func NewGetEventsForMonthUseCase(
	eventRepository domain.Repository,
) *GetEventsForMonthUseCase {
	return &GetEventsForMonthUseCase{
		eventRepository: eventRepository,
	}
}

// Execute отвечает за выполнение бизнес-логики для операции получения события по месяцу
func (uc *GetEventsForMonthUseCase) Execute(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return uc.eventRepository.GetEventsForMonth(ctx, date)
}
