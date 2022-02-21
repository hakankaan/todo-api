package postgres

import (
	"fmt"
	"os"

	"github.com/hakankaan/todo-api/pkg/config"
	"github.com/hakankaan/todo-api/pkg/logging"
	"github.com/hakankaan/todo-api/pkg/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// service is a struct for database access service
type service struct {
	l      logging.Service
	DB     *gorm.DB
	Config config.Service
}

// New constructor of the database access service
func New(l logging.Service, c config.Service) (s *service) {
	h := c.CFG.Postgres.Host
	p := c.CFG.Postgres.Port
	u := c.CFG.Postgres.User
	dbname := c.CFG.Postgres.DB
	pass := c.CFG.Postgres.Pass
	sslmode := c.CFG.Postgres.SSL

	db, err := newConnection(h, p, u, dbname, pass, sslmode)
	if err != nil {
		l.Error("newConnection", err)
		os.Exit(1)
	}

	s = &service{
		l:  l,
		DB: db,
	}

	err = s.migrateAll()
	if err != nil {
		l.Error("migrateAll", err)
		os.Exit(1)
	}

	return
}

// newConnection creates a new connection to the database
func newConnection(h, p, u, dbname, pass, sslmode string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", h, p, u, dbname, pass, sslmode)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

// migrateAll migrates all tables
func (s *service) migrateAll() (err error) {
	err = s.DB.AutoMigrate(&store.Todo{})
	return
}
