package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"project/service1/config"
	"time"
)

// Timeout is a Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*gorm.DB
}

// Dial creates new database connection to postgres
func Dial() (*DB, error) {
	cfg := config.Get()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Etc/GMT-3",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Port,
	)

	var db *gorm.DB
	var err error

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	/*postgres.Open()
	pgOpts := &pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Database.Host, cfg.Database.Port),
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Database: cfg.Database.DbName,
	}*/

	//pgDB := pg.Connect(pgOpts)

	// run test select query to make sure PostgreSQL is up and running
	var attempt uint

	const maxAttempts = 10

	for {
		attempt++

		log.Printf("[PostgreSQL.Dial] (Ping attempt %d) SELECT 1\n", attempt)

		//_, err := pgDB.Exec("SELECT 1")
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			log.Printf("[PostgreSQL.Dial] (Ping attempt %d) error: %s\n", attempt, err)

			if attempt < maxAttempts {
				time.Sleep(1 * time.Second)

				continue
			}

			return nil, fmt.Errorf("pgDB.Exec failed: %w", err)
		}

		log.Printf("[PostgreSQL.Dial] (Ping attempt %d) OK\n", attempt)

		break
	}

	sqlDB, _ := db.DB()

	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(Timeout))
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return &DB{db}, nil
}
