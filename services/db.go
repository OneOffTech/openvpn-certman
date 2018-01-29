package services

import (
	"errors"
	"log"

	"git.klink.asia/paul/certman/models"
	"git.klink.asia/paul/certman/settings"
	"github.com/jinzhu/gorm"
)

// Error Definitions
var (
	ErrNotImplemented = errors.New("Not implemented")
)

var Database *gorm.DB

// DB is a wrapper around gorm.DB to provide custom methods
type DB struct {
	*gorm.DB
}

func InitDB() *gorm.DB {
	dsn := settings.Get("DATABASE_URL", "db.sqlite3")

	// Establish connection
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Could not open database: %s", err.Error())
	}

	// Migrate models
	db.AutoMigrate(models.User{}, models.Client{})
	db.LogMode(true)

	Database = db
	return db
}

// CountUsers returns the number of Users in the datastore
func (db *DB) CountUsers() (uint, error) {
	return 0, ErrNotImplemented
}

// CreateUser inserts a user into the datastore
func (db *DB) CreateUser(*models.User) (*models.User, error) {
	return nil, ErrNotImplemented
}

// ListUsers returns a slice of 'count' users, starting at 'offset'
func (db *DB) ListUsers(count, offset int) ([]*models.User, error) {
	var users = make([]*models.User, 0)

	return users, ErrNotImplemented
}

// GetUserByID returns a single user by ID
func (db *DB) GetUserByID(id uint) (*models.User, error) {
	return nil, ErrNotImplemented
}

// GetUserByEmail returns a single user by email
func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	return nil, ErrNotImplemented
}

// DeleteUser removes a user from the datastore
func (db *DB) DeleteUser(id uint) error {
	return ErrNotImplemented
}
