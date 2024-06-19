package ports

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"webtech/http/internal/calendar/builder"
	"webtech/http/internal/calendar/domain"
)

// HTTPCalendarHandler представляет собой обработчика REST-операций
type HTTPCalendarHandler struct {
	app *builder.Application
}

// NewHTTPCalendarHandler отвечает за создание нового оброботчика REST-операций
func NewHTTPCalendarHandler(app *builder.Application) HTTPCalendarHandler {
	return HTTPCalendarHandler{app: app}
}

// CreateEvent отвечает за REST-операцию по созданию события
func (h HTTPCalendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.mapToResponse(w, http.StatusMethodNotAllowed, nil, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	event, statusCode, errMessage := h.validationAndParse(r)
	if statusCode != 200 {
		h.mapToResponse(w, statusCode, nil, errMessage)
		return
	}

	if event.UserID == 0 || event.Date == (time.Time{}) || event.Description == "" {
		h.mapToResponse(w, http.StatusBadRequest, nil, http.StatusText(http.StatusBadRequest))
		return
	}

	id, err := h.app.CreateEvent.Execute(r.Context(), event)
	if err != nil {
		h.mapToResponse(w, http.StatusServiceUnavailable, nil, err.Error())
		return
	}

	event.ID = id

	h.mapToResponse(w, http.StatusOK, event, "")
}

// UpdateEvent отвечает за REST-операцию по обновлению события
func (h HTTPCalendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.mapToResponse(w, http.StatusMethodNotAllowed, nil, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	event, statusCode, errMessage := h.validationAndParse(r)
	if statusCode != 200 {
		h.mapToResponse(w, statusCode, nil, errMessage)
		return
	}

	if event.ID == 0 || event.UserID == 0 || event.Date == (time.Time{}) || event.Description == "" {
		h.mapToResponse(w, http.StatusBadRequest, nil, http.StatusText(http.StatusBadRequest))
		return
	}

	err := h.app.UpdateEvent.Execute(r.Context(), event)
	if err != nil {
		h.mapToResponse(w, http.StatusServiceUnavailable, nil, err.Error())
		return
	}

	h.mapToResponse(w, http.StatusOK, event, "")
}

// DeleteEvent отвечает за REST-операцию по удалению события
func (h HTTPCalendarHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.mapToResponse(w, http.StatusMethodNotAllowed, nil, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	event, statusCode, errMessage := h.validationAndParse(r)
	if statusCode != 200 {
		h.mapToResponse(w, statusCode, nil, errMessage)
		return
	}

	if event.ID == 0 {
		h.mapToResponse(w, http.StatusBadRequest, nil, http.StatusText(http.StatusBadRequest))
		return
	}

	event, err := h.app.GetEventByID.Execute(r.Context(), event.ID)
	if err != nil {
		h.mapToResponse(w, http.StatusServiceUnavailable, nil, err.Error())
		return
	}

	err = h.app.DeleteEvent.Execute(r.Context(), event.ID)
	if err != nil {
		h.mapToResponse(w, http.StatusServiceUnavailable, nil, err.Error())
		return
	}

	h.mapToResponse(w, http.StatusOK, event, "")
}

// GetEventsForDay отвечает за REST-операцию по получению событий по дню
func (h HTTPCalendarHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.mapToResponse(w, http.StatusMethodNotAllowed, nil, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	event, statusCode, errMessage := h.validationAndParse(r)
	if statusCode != 200 {
		h.mapToResponse(w, statusCode, nil, errMessage)
		return
	}

	if event.Date == (time.Time{}) {
		h.mapToResponse(w, http.StatusBadRequest, nil, http.StatusText(http.StatusBadRequest))
		return
	}

	events, err := h.app.GetEventsForDay.Execute(r.Context(), event.Date)
	if err != nil {
		h.mapToResponse(w, http.StatusServiceUnavailable, nil, err.Error())
		return
	}

	h.mapToResponse(w, http.StatusOK, events, "")
}

// GetEventsForWeek отвечает за REST-операцию по получению событий по неделе
func (h HTTPCalendarHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.mapToResponse(w, http.StatusMethodNotAllowed, nil, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	event, statusCode, errMessage := h.validationAndParse(r)
	if statusCode != 200 {
		h.mapToResponse(w, statusCode, nil, errMessage)
		return
	}

	if event.Date == (time.Time{}) {
		h.mapToResponse(w, http.StatusBadRequest, nil, http.StatusText(http.StatusBadRequest))
		return
	}

	events, err := h.app.GetEventsForWeek.Execute(r.Context(), event.Date)
	if err != nil {
		h.mapToResponse(w, http.StatusServiceUnavailable, nil, err.Error())
		return
	}

	h.mapToResponse(w, http.StatusOK, events, "")
}

// GetEventsForMonth отвечает за REST-операцию по получению событий по месяцу
func (h HTTPCalendarHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.mapToResponse(w, http.StatusMethodNotAllowed, nil, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	event, statusCode, errMessage := h.validationAndParse(r)
	if statusCode != 200 {
		h.mapToResponse(w, statusCode, nil, errMessage)
		return
	}

	if event.Date == (time.Time{}) {
		h.mapToResponse(w, http.StatusBadRequest, nil, http.StatusText(http.StatusBadRequest))
		return
	}

	events, err := h.app.GetEventsForMonth.Execute(r.Context(), event.Date)
	if err != nil {
		h.mapToResponse(w, http.StatusServiceUnavailable, nil, err.Error())
		return
	}

	h.mapToResponse(w, http.StatusOK, events, "")
}

// MiddlewareLogger отвечает за логгирование запросов
func (h HTTPCalendarHandler) MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s\n", r.Method, r.RequestURI)
		next(w, r)
	}
}

// CustomRegisterHandlers отвечает за инициализацию ручек роутера
func CustomRegisterHandlers(router *http.ServeMux, h HTTPCalendarHandler) {
	router.HandleFunc("/create_event", h.MiddlewareLogger(h.CreateEvent))
	router.HandleFunc("/update_event", h.MiddlewareLogger(h.UpdateEvent))
	router.HandleFunc("/delete_event", h.MiddlewareLogger(h.DeleteEvent))
	router.HandleFunc("/events_for_day", h.MiddlewareLogger(h.GetEventsForDay))
	router.HandleFunc("/events_for_week", h.MiddlewareLogger(h.GetEventsForWeek))
	router.HandleFunc("/events_for_month", h.MiddlewareLogger(h.GetEventsForMonth))
}

type jsonEvent struct {
	ID          int       `json:"event_id"`
	UserID      int       `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

func (h HTTPCalendarHandler) validationAndParse(r *http.Request) (domain.Event, int, string) {
	event := domain.Event{}

	if r.Method == http.MethodGet || (r.Method == http.MethodPost && r.Header.Get("Content-Type") == "application/x-www-form-urlencoded") {
		err := r.ParseForm()
		if err != nil {
			return domain.Event{}, http.StatusBadRequest, err.Error()
		}

		if r.Form.Get("event_id") != "" {
			event.ID, err = strconv.Atoi(r.Form.Get("event_id"))
			if err != nil {
				return domain.Event{}, http.StatusBadRequest, err.Error()
			}
		}

		if r.Form.Get("user_id") != "" {
			event.UserID, err = strconv.Atoi(r.Form.Get("user_id"))
			if err != nil {
				return domain.Event{}, http.StatusBadRequest, err.Error()
			}
		}

		if r.Form.Get("date") != "" {
			event.Date, err = time.Parse("2006-01-02", r.Form.Get("date"))
			if err != nil {
				return domain.Event{}, http.StatusBadRequest, err.Error()
			}
		}

		event.Description = r.Form.Get("description")
	} else if r.Header.Get("Content-Type") == "application/json" {
		jEvent := jsonEvent{}

		err := json.NewDecoder(r.Body).Decode(&jEvent)
		if err != nil {
			return domain.Event{}, http.StatusBadRequest, err.Error()
		}

		event.ID = jEvent.ID
		event.UserID = jEvent.UserID
		event.Date = jEvent.Date
		event.Description = jEvent.Description
	} else {
		return domain.Event{}, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)
	}

	return event, http.StatusOK, ""
}

func (h HTTPCalendarHandler) mapToResponse(w http.ResponseWriter, statusCode int, data interface{}, errMessage string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	response := make(map[string]interface{})

	if statusCode >= 200 && statusCode < 300 {
		response["result"] = data
	} else {
		response["error"] = errMessage
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
