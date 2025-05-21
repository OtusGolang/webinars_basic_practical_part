package app

import (
	"context"
	"fmt"
	"gohttp2/internal/config"
	"gohttp2/internal/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	cfg *config.Config
	ctx context.Context
}

func NewService(ctx context.Context) (*App, error) {
	// Инит баз клиентов
	return &App{
		ctx: ctx,
		cfg: config.NewConfig(),
	}, nil
}

func (a *App) Start() error {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	mux := http.NewServeMux()

	// mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	handler.Base(a.ctx, w, r)
	// })
	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		handler.Login(a.ctx, w, r)
	})
	mux.HandleFunc("PUT /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.UserUpdate(a.ctx, w, r)
	})

	serverHTTP := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler: mux,
	}

	go func() {
		log.Println("server starting at ", serverHTTP.Addr)
		if err := serverHTTP.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	ctxT, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := serverHTTP.Shutdown(ctxT); err != nil {
		return fmt.Errorf("shutdown server: %s\n", err)
	}
	log.Println("FINAL server shutdown")
	return nil
}
