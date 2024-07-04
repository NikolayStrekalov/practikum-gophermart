package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/NikolayStrekalov/practicum-gophermart/internal/config"
)

func Start(ctx context.Context) {
	// настройка роутов
	r := NewRouter()

	// Создаем сервер и запускаем его
	var server *http.Server
	go func() {
		<-ctx.Done()
		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()
	server = &http.Server{Addr: config.AppConfig.SelfAddress, Handler: r}
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
