package usecase

import (
	"context"
	"webtech/http/internal/calendar/domain"
)

// DeleteEventUseCase представляет обработчик бизнес-логики для операции удаления события
type DeleteEventUseCase struct {
	eventRepository domain.Repository
}

// NewDeleteEventUseCase отвечает за инициализацию обработчика бизнес-логики для операции удаления события
func NewDeleteEventUseCase(eventRepository domain.Repository) *DeleteEventUseCase {
	return &DeleteEventUseCase{
		eventRepository: eventRepository,
	}
}

// Execute отвечает за выполнение бизнес-логики для операции удаления события
func (uc *DeleteEventUseCase) Execute(ctx context.Context, eventID int) error {
	return uc.eventRepository.DeleteEvent(ctx, eventID)
}
