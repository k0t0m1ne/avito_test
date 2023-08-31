package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprint(resp, m)
	if err != nil {
		return
	}
}

func Serve(ctx context.Context) {

	// Задаем параметры сервера и пути
	server := http.Server{Addr: ":8080"}
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userSegmentHandler).Methods(http.MethodGet)
	router.HandleFunc("/segments/add", addSegmentHandler).Methods(http.MethodPost)
	router.HandleFunc("/segments/delete", deleteSegmentHandler).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", addUserToSegment).Methods(http.MethodPut)
	http.Handle("/", router)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()
	// Проверяем, что у нас не подан сигнал на завершение работы
	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down server")
			err := server.Shutdown(ctx)
			if err != nil {
				panic(err)
			}
			return
		}
	}
}
