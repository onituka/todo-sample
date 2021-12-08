package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/onituka/todo-sample/infrastructure/persistence"
	"github.com/onituka/todo-sample/infrastructure/persistence/rdb"
	"github.com/onituka/todo-sample/interfaces/handler"
	"github.com/onituka/todo-sample/usecase"
)

func Run() error {
	mySQLHandler, err := rdb.NewMySQLHandler()
	if err != nil {
		return err
	}
	defer mySQLHandler.Conn.Close()

	todoRepository := persistence.NewTodoRepository(mySQLHandler)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	router := mux.NewRouter()

	router.HandleFunc("/todos/{id:[0-9]+}", todoHandler.FetchTodo).Methods(http.MethodGet)
	router.HandleFunc("/todos", todoHandler.FetchAllTodo).Methods(http.MethodGet)
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id:[0-9]+}", todoHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id:[0-9]+}", todoHandler.Delete).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8081),
		Handler: router,
	}

	errorCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errorCh <- err
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errorCh:
		return err
	case s := <-signalCh:
		log.Printf("SIGNAL %s received", s.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}
