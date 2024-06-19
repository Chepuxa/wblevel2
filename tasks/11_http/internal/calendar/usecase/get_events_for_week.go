package usecase

import (
	"context"
	"time"
	"webtech/http/internal/calendar/domain"
)

// GetEventsForWeekUseCase представляет обработчик бизнес-логики для операции получения события по неделе
type GetEventsForWeekUseCase struct {
	eventRepository domain.Repository
}

// NewGetEventsForWeekUseCase отвечает за инициализацию обработчика бизнес-логики для операции получения события по неделе
func NewGetEventsForWeekUseCase(
	eventRepository domain.Repository,
) *GetEventsForWeekUseCase {
	return &GetEventsForWeekUseCase{
		eventRepository: eventRepository,
	}
}

// Execute отвечает за выполнение бизнес-логики для операции получения события по месяцу
func (uc *GetEventsForWeekUseCase) Execute(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return uc.eventRepository.GetEventsForWeek(ctx, date)
}
