package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/config"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/db"
	logEntity "github.com/kissejau/backend-trainee-assignment-2023/internal/server/log"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/logger"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/segment"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/user"
)

type Server interface {
	Run()
}

type server struct {
	port   string
	r      *httprouter.Router
	db     *sql.DB
	logger http.Handler
}

func NewServer() Server {
	cfg := config.NewConfig()

	port := cfg.Port

	db, err := db.NewDb(cfg.Db)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}

	r := httprouter.New()

	logger := logger.Logger(r)

	return &server{
		port:   port,
		r:      r,
		db:     db,
		logger: logger,
	}
}

func (s *server) Run() {
	s.Startup()
	http.ListenAndServe(":8080", s.logger)
}

func (s *server) Startup() {
	segmentRepo := segment.NewRepository(s.db)
	userRepo := user.NewRepository(s.db, segmentRepo)
	historyRepo := logEntity.NewRepository(s.db)

	segmentHandler := segment.NewHandler(segmentRepo)
	userHandler := user.NewHandler(userRepo)
	historyHandler := logEntity.NewHandler(historyRepo)

	userHandler.Register(s.r)
	segmentHandler.Register(s.r)
	historyHandler.Register(s.r)
}
