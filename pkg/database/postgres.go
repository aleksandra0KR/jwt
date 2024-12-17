package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jwt/domain"
	"os"
	"time"
)

type Postgres struct {
	db                 *gorm.DB
	MaxIdleConnections int
	MaxOpenConnections int
}

func InitializeDBPostgres(maxIdleConnections, maxOpenConnections int) *Postgres {
	postgresDB := Postgres{
		MaxIdleConnections: maxIdleConnections,
		MaxOpenConnections: maxOpenConnections,
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connectionDBUrl := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s`, dbHost, dbUser, dbPassword, dbName, dbPort)
	log.Infof(connectionDBUrl)

	var db *gorm.DB
	var err error
	maxRetries := 10
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(connectionDBUrl), &gorm.Config{})
		if err == nil {
			break
		}
		log.Warnf("failed to connect to database: %v. Retrying in %v...", err, retryDelay)
		time.Sleep(retryDelay)
	}
	if err != nil {
		log.Fatalf("failed to connect to database after %d retries: %v", maxRetries, err)
	}
	postgresDB.db = db

	sqlDB, err := postgresDB.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(postgresDB.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(postgresDB.MaxOpenConnections)

	postgresDB.db = db
	log.Info("connected to Postgres DB")

	postgresDB.Migrate()
	return &postgresDB
}

func (postgresDB *Postgres) Migrate() {
	if err := postgresDB.db.Migrator().DropTable(&domain.User{}, &domain.RefreshToken{}); err != nil {
		log.Fatal("failed to drop tables:", err)
	}

	if err := postgresDB.db.AutoMigrate(&domain.User{}, &domain.RefreshToken{}); err != nil {
		log.Fatal("failed to create tables:", err)
	}

	user := domain.User{
		Guid:  "123",
		Email: "example@gmail.com",
		IP:    "456789",
	}
	if err := postgresDB.db.Create(&user).Error; err != nil {
		log.Fatal("failed to insert data into users table:", err)
	}
}

func (postgresDB *Postgres) GetDB() *gorm.DB {
	return postgresDB.db
}
