package dependencies

import (
	"book-api/internal/config"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/mysqldialect"

	_ "github.com/go-sql-driver/mysql"
)

type Dependencies struct {
	DB *bun.DB
}

func InitDependencies(cfg *config.Config) (*Dependencies, error) {
	db, err := InitDB(cfg.MySQL.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	return &Dependencies{DB: db}, nil
}

func InitDB(dsn string) (*bun.DB, error) {
	fmt.Println(dsn)

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return bun.NewDB(sqldb, mysqldialect.New()), nil
}
