package databases

import (
	"database/sql"
	"fmt"

	configs "github.com/kittiphong064008-cyber/Go-Pokemon/config"
	_ "github.com/lib/pq"
)

func ConnectPostgres(config configs.Configs) (*sql.DB, error) {
	cfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.PostgreSQL.Host,
		config.PostgreSQL.Port,
		config.PostgreSQL.User,
		config.PostgreSQL.Password,
		config.PostgreSQL.Database,
		config.PostgreSQL.SSLMode)
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
