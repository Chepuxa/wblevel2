package usecase

import (
	"context"
	"time"
	"webtech/http/internal/calendar/domain"
)

// GetEventsForDayUseCase представляет обработчик бизнес-логики для операции получения события по дню
type GetEventsForDayUseCase struct {
	eventRepository domain.Repository
}

// NewGetEventsForDayUseCase отвечает за инициализацию обработчика бизнес-логики для операции получения события по дню
func NewGetEventsForDayUseCase(
	eventRepository domain.Repository,
) *GetEventsForDayUseCase {
	return &GetEventsForDayUseCase{
		eventRepository: eventRepository,
	}
}

// Execute отвечает за выполнение бизнес-логики для операции получения события по дню
func (uc *GetEventsForDayUseCase) Execute(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return uc.eventRepository.GetEventsForDay(ctx, date)
}
