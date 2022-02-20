package postgres

import (
	"os"

	"github.com/hakankaan/todo-api/pkg/logging"
	"github.com/hakankaan/todo-api/pkg/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// service is a struct for database access service
type service struct {
	l  logging.Service
	DB *gorm.DB
}

// New constructor of the database access service
func New(l logging.Service) (s *service, err error) {
	db, err := newConnection()
	if err != nil {
		os.Exit(1)
	}

	s = &service{
		l:  l,
		DB: db,
	}

	err = s.migrateAll()
	if err != nil {
		os.Exit(1)
	}

	return
}

// newConnection creates a new connection to the database
func newConnection() (db *gorm.DB, err error) {
	dsn := "host=postgres port=5432 user=postgres dbname=todo password=postgres sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

// migrateAll migrates all tables
func (s *service) migrateAll() (err error) {
	err = s.DB.AutoMigrate(&store.Todo{})
	return
}
