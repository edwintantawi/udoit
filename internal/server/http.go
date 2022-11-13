package server

import (
	"database/sql"
	"net/http"

	"github.com/edwintantawi/udoit/internal/task"
	"github.com/edwintantawi/udoit/pkg/idgen"
	"github.com/gorilla/mux"
)

type httpServer struct {
	db *sql.DB
}

func NewHTTP(db *sql.DB) *httpServer {
	return &httpServer{db: db}
}

func (s *httpServer) Setup(addr string) *http.Server {
	mux := mux.NewRouter()

	idgen := idgen.New()

	taskRepo := task.NewRepo(s.db, idgen)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)

	task.NewRouter("/tasks", mux, taskHandler)

	return &http.Server{Addr: addr, Handler: mux}
}
