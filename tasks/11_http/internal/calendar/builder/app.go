package builder

import (
	"context"
	"webtech/http/internal/calendar/adapters"
	"webtech/http/internal/calendar/usecase"
)

// Application представляет структуру, которая хранит все операции бизнес-логики
type Application struct {
	CreateEvent       *usecase.CreateEventUseCase
	UpdateEvent       *usecase.UpdateEventUseCase
	DeleteEvent       *usecase.DeleteEventUseCase
	GetEventByID      *usecase.GetEventByIDUseCase
	GetEventsForDay   *usecase.GetEventsForDayUseCase
	GetEventsForWeek  *usecase.GetEventsForWeekUseCase
	GetEventsForMonth *usecase.GetEventsForMonthUseCase
}

// NewApplication отвечает за инициализацию Application
func NewApplication(ctx context.Context) *Application {
	eventRepository := adapters.NewCacheEventRepository(200)

	return &Application{
		CreateEvent:       usecase.NewCreateEventUseCase(eventRepository),
		UpdateEvent:       usecase.NewUpdateEventUseCase(eventRepository),
		DeleteEvent:       usecase.NewDeleteEventUseCase(eventRepository),
		GetEventByID:      usecase.NewGetEventByIDUseCase(eventRepository),
		GetEventsForDay:   usecase.NewGetEventsForDayUseCase(eventRepository),
		GetEventsForWeek:  usecase.NewGetEventsForWeekUseCase(eventRepository),
		GetEventsForMonth: usecase.NewGetEventsForMonthUseCase(eventRepository),
	}
}
