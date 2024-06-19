package usecase

import (
	"context"
	"webtech/http/internal/calendar/domain"
)

// UpdateEventUseCase представляет обработчик бизнес-логики для операции обновления события
type UpdateEventUseCase struct {
	eventRepository domain.Repository
}

// NewUpdateEventUseCase отвечает за инициализацию обработчика бизнес-логики для операции обновления события
func NewUpdateEventUseCase(
	eventRepository domain.Repository,
) *UpdateEventUseCase {
	return &UpdateEventUseCase{
		eventRepository: eventRepository,
	}
}

// Execute отвечает за выполнение бизнес-логики для операции обновления события
func (uc *UpdateEventUseCase) Execute(ctx context.Context, updatedEvent domain.Event) error {
	return uc.eventRepository.UpdateEvent(ctx, updatedEvent)
}
