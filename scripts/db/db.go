package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	GetDB() *sql.DB
}

type database struct {
	db *sql.DB
}

func NewDatabase() Database {
	db, err := connectDB()
	if err != nil {
		log.Printf("error:%s", err.Error())
		panic(err)
	}
	return &database{db}
}

func (d *database) GetDB() *sql.DB {
	return d.db
}

func connectDB() (*sql.DB, error) {
	conf, err := getDBConfig(os.Getenv("ENV"))
	if err != nil {
		return nil, err
	}
	dbConf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
	db, err := sql.Open("mysql", dbConf)
	if err != nil {
		return db, err
	}

	// lambdaは、無限ループだと危険なため、リトライは5回で制限
	maxAttempts := 5
	for i := 0; i < maxAttempts; i++ {
		if err = db.Ping(); err != nil {
			log.Print(err.Error())
			if i < maxAttempts-1 {
				time.Sleep(10 * time.Second)
				continue
			}
		} else {
			return db, nil
		}
	}
	return db, err
}

type DBConfig struct {
	DBName     string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
}

func getDBConfig(env string) (DBConfig, error) {
	conf := DBConfig{}
	switch env {
	case "local":
		// ローカル用に立ち上げたコンテナDBの接続情報であるためハードコーディング
		conf.DBName = "db_name"
		conf.DBHost = "mysql"
		conf.DBPort = 3306
		conf.DBUser = "user"
		conf.DBPassword = "password"
		return conf, nil
	default:
		return conf, errors.New("failed to get db conf")
	}
}
