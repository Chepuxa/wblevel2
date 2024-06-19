package usecase

import (
	"context"
	"webtech/http/internal/calendar/domain"
)

// CreateEventUseCase представляет обработчик бизнес-логики для операции создания события
type CreateEventUseCase struct {
	eventRepository domain.Repository
}

// NewCreateEventUseCase отвечает за инициализацию обработчика бизнес-логики для операции создания события
func NewCreateEventUseCase(eventRepository domain.Repository) *CreateEventUseCase {
	return &CreateEventUseCase{
		eventRepository: eventRepository,
	}
}

// Execute отвечает за выполнение бизнес-логики для операции создания события
func (uc *CreateEventUseCase) Execute(ctx context.Context, event domain.Event) (int, error) {
	return uc.eventRepository.CreateEvent(ctx, event)
}
