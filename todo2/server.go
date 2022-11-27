package todo

// для запуска http  сервера
import (
	"context"
	"net/http"
	"time"
)

type Server struct { // структура сервер для запуска http server
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port, // значение порта на котором будет запускаться сервер
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1MB // запуск
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe() // метод станд http servera
	// под капотом запускает цикл бесконечный и слушает все входящие запросы по
	// обработке
}

func (s *Server) Shutdown(ctx context.Context) error { // для выхода их приложения
	return s.httpServer.Shutdown(ctx)
}
