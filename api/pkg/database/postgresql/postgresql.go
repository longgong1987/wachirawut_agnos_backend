package postgresql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
)

type Config struct {
	Host     string `env:"POSTGRESQL_HOST"`
	Username string `env:"POSTGRESQL_USERNAME"`
	Password string `env:"POSTGRESQL_PASSWORD"`
	Database string `env:"POSTGRESQL_DATABASE"`
	Port     int    `env:"POSTGRESQL_PORT"`
}

func buildDSN(config *Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=UTC",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)
}

func ConnectDB(dbConf *Config) *gorm.DB {
	dsn := buildDSN(dbConf)

	logger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(`fatal error: cannot connect to database`)
	}

	log.Println("PostgreSQL connected successfully!!!")

	return dbConn
}
