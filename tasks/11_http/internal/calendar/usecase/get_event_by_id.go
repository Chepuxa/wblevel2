package usecase

import (
	"context"
	"webtech/http/internal/calendar/domain"
)

// GetEventByIDUseCase представляет обработчик бизнес-логики для операции получения события по ID
type GetEventByIDUseCase struct {
	eventRepository domain.Repository
}

// NewGetEventByIDUseCase отвечает за инициализацию обработчика бизнес-логики для операции получения события по ID
func NewGetEventByIDUseCase(
	eventRepository domain.Repository,
) *GetEventByIDUseCase {
	return &GetEventByIDUseCase{
		eventRepository: eventRepository,
	}
}

// Execute отвечает за выполнение бизнес-логики для операции получения события по ID
func (uc *GetEventByIDUseCase) Execute(ctx context.Context, eventID int) (domain.Event, error) {
	return uc.eventRepository.GetEventByID(ctx, eventID)
}
