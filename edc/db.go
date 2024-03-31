package edc

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBSetup(params DBSetupArgs) *gorm.DB {
	log.Info("Setting up DB connection...")
	var dbOpen gorm.Dialector

	switch params.Engine {
	case SQLite:
		dbOpen = SQLiteConn(params.DSN)
	case Postgres:
		dbOpen = PostgresConn(params.DSN)
	case MySQL:
		dbOpen = MySQLConn(params.DSN)
	default:
		msg := fmt.Sprintf(
			"Invalid DB Engine. Valid options are: %s, %s, %s. You provided: %s",
			SQLite, Postgres, MySQL, params.Engine,
		)
		panic(msg)
	}

	db, err := gorm.Open(dbOpen, &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Failed to connect to DB: %s", err)
		panic(msg)
	}

	log.Info("DB connection established.")
	return db
}

func SQLiteConn(params DBDSN) gorm.Dialector {
	dsn := fmt.Sprintf("%s.db", params.Database)
	return sqlite.Open(dsn)
}

func PostgresConn(params DBDSN) gorm.Dialector {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		params.Host,
		params.User,
		params.Password,
		params.Database,
		params.Port)
	return postgres.Open(dsn)
}

func MySQLConn(params DBDSN) gorm.Dialector {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.Database)
	return mysql.Open(dsn)
}
