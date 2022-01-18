package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/config"
	"github.com/pkg/errors"
)

type Article struct {
	conn *sql.DB
}

const DbConnectTimeoutSecs int = 15
const DbExecTimeoutSecs int = 15

func getConnection(schema string, cfg *config.Database) (*sql.DB, error) {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&timeout=%ds",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		schema,
		DbConnectTimeoutSecs)

	db, err := sql.Open("mysql", str)

	if err != nil {
		return nil, fmt.Errorf("error establishing a connection: %v, conn details: %v", err.Error(), str)
	}

	db.SetConnMaxLifetime(time.Second * time.Duration(DbExecTimeoutSecs))
	return db, nil
}

func NewStore(cfg *config.Database) (*Article, error) {
	if cfg == nil {
		return nil, fmt.Errorf("missing database cfg parameter")
	}

	db, err := getConnection(cfg.Schema, cfg)

	if err != nil {
		return nil, errors.Wrap(err, "error trying to establish a connection")
	}

	return &Article{
		conn: db,
	}, nil

}
