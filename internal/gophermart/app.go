package gophermart

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
)

func Start() {
	config := GetConfig()

	// настройка роутов
	r := NewRouter()

	// Создаем сервер и запускаем его
	var server *http.Server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()
	server = &http.Server{Addr: config.SelfAddress, Handler: r}
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
