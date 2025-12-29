package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yporn/doh-ems-api/config"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
	db  *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) Intialize() {

} 