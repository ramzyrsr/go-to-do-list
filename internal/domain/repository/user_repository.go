package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"to-do-list/internal/domain/models"
	"to-do-list/internal/infrastructure/middleware"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Login(user *models.Login) (*models.Login, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer tx.Rollback() // Ensure the transaction is rolled back if there’s an error

	var existingEmail string
	checkQuery := "SELECT email FROM users WHERE email = $1"
	err = r.db.QueryRow(checkQuery, user.Email).Scan(&existingEmail)

	if err == nil {
		// Email already exists, return an error
		return nil, fmt.Errorf("email '%s' already exists. Please use a different email", user.Email)
	} else if err != sql.ErrNoRows {
		// Handle other errors (e.g., query error)
		return nil, fmt.Errorf("failed to check email existence: %v", err)
	}

	hashPass, err := middleware.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	query := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING uuid"
	err = r.db.QueryRow(query, user.Name, user.Email, hashPass).Scan(&user.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user into database: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return user, nil
}

func (r *userRepository) Login(user *models.Login) (*models.Login, error) {
	start := time.Now()
	pass := user.Password

	// Prepare the query using Prepare(), this ensures the query is parsed and compiled once
	query := "SELECT uuid, email, password FROM users WHERE email = $1"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %v", err)
	}
	defer stmt.Close() // Ensure the statement is closed when done
	ttfb := time.Since(start)
	log.Printf("Request processed in %s (TTFB) 1", ttfb)

	// Execute the prepared statement with the email as parameter
	err = stmt.QueryRow(user.Email).Scan(&user.UUID, &user.Email, &user.Password)
	fmt.Println(err, user.Email, user.Password, pass)
	ttfb = time.Since(start)
	log.Printf("Request processed in %s (TTFB) 2", ttfb)
	// Check for database error
	if err == sql.ErrNoRows {
		// No user found with this email, return a generic error to avoid leaking info
		return nil, fmt.Errorf("email or password is not match. please check again your data")
	}
	if err != nil {
		// If there was any other database error, log it and return a generic error
		return nil, fmt.Errorf("something went wrong, please try again later")
	}
	ttfb = time.Since(start)
	log.Printf("Request processed in %s (TTFB) 3", ttfb)

	isMatch := middleware.CheckPasswordHash(pass, user.Password)
	ttfb = time.Since(start)
	log.Printf("Request processed in %s (TTFB) 4", ttfb)
	fmt.Println(isMatch)
	if !isMatch {
		return nil, fmt.Errorf("email or password is not match. please check again your data")
	}
	ttfb = time.Since(start)
	log.Printf("Request processed in %s (TTFB) 5", ttfb)

	return user, nil
}
