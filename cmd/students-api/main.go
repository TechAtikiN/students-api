package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/techatikin/students-api/internal/config"
	"github.com/techatikin/students-api/internal/http/handlers/student"
	"github.com/techatikin/students-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("Error creating database:", slog.String("error", err.Error()))
	}

	slog.Info("Storage initialized")

	// router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("PUT /api/students/{id}", student.UpdateById(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteById(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started %s", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Error starting server:", err)
		}
	}()

	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Error shutting down server:", slog.String("error", err.Error()))
	}

	slog.Info("Server stopped")
	slog.Info("Bye!")
}
