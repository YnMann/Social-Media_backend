package store

import (
	"database/sql"
	"fmt"
	"time"

	// Импортируем анонимно, чтобы методы не импортировались
	_ "github.com/lib/pq"
)

// Store
type Store struct {
	config *Config
	db     *sql.DB
}

// New
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	// Устанавливаем таймаут на 10 секунд для попытки установления соединения
	db.SetConnMaxLifetime(10 * time.Second)

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	fmt.Println("Успешное подключение к базе данных!")

	return nil
}

// Close
func (s *Store) Close() {
	s.db.Close()
}
