package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"simple-api-go2/api/utils"
	"time"

	"github.com/err-him/gonf"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host   string
	Port   string
	Uname  string
	DBName string
	Pass   string
}

type DB struct {
	SQL *sql.DB
}

var dbCon = &DB{}

func ConnectDB() (*DB, error) {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}
	cfg := DBConfig{}
	err := gonf.GetConfig(utils.GetEnvDBFile(env), &cfg)
	if err != nil {
		log.Fatal("DB Details can not be loaded, shutting down the application")
	}

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Uname,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(16)
	db.SetConnMaxLifetime(30 * time.Minute)
	dbCon.SQL = db
	return dbCon, err
}
