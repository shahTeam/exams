package connection

import (
	"database/sql"
	//"errors"
	"exam/config"
	"fmt"

	//"github.com/golang-migrate/migrate/v4"
	//"github.com/golang-migrate/migrate/v4/database/postgres"
	//"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
		),
	)
	if err != nil {
		return nil, err
	}

	// if err := db.Ping(); err != nil {
	// 	return nil, err
	// }

	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	return nil, err
	// }

	// m, err := migrate.NewWithDatabaseInstance("file://repository/migrations", "postgres", driver)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange){
	// 	return nil, err
	// }


	return db, nil
}
