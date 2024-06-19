package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webtech/http/internal/calendar/builder"
	"webtech/http/internal/calendar/ports"
)

type app struct {
	httpServer *http.Server
}

func (a *app) run(addr string) error {
	router := http.NewServeMux()

	ctx := context.Background()

	calendarApp := builder.NewApplication(ctx)
	calendarHTTPHandler := ports.NewHTTPCalendarHandler(calendarApp)
	ports.CustomRegisterHandlers(router, calendarHTTPHandler)

	a.httpServer = &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    180 * time.Second,
		WriteTimeout:   180 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server is running...")

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	doneCh := make(chan struct{})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	close(doneCh)

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func main() {
	app := &app{}
	err := app.run(":8080")
	if err != nil {
		panic(err)
	}
}
