package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gchaincl/dotsql"
	_ "github.com/lib/pq"
)

// Conn contains the queries and the db connection pool
type Conn struct {
	Queries *dotsql.DotSql
	db      *sql.DB
}

// Config ...
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Username string `json:"-"`
	Password string `json:"-"`
}

// New returns a pg connection pool
func New(ctx context.Context, c Config) (*Conn, error) {
	files, err := ioutil.ReadDir("./store/sql")
	if err != nil {
		return nil, err
	}

	dots := make([]*dotsql.DotSql, 0, len(files))
	for _, f := range files {
		dot, err := dotsql.LoadFromFile(filepath.Join("./store/sql", f.Name()))
		if err != nil {
			return nil, err
		}
		dots = append(dots, dot)
	}
	queries := dotsql.Merge(dots[0])
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s, dbname=%s sslmode=disable", "localhost", "5432", c.Username, c.Password, c.Name))
	if err != nil {
		return nil, err
	}

	res := &Conn{Queries: queries, db: db}

	return res, db.PingContext(ctx)
}

// QueryContext ...
func (c *Conn) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.QueryContext(ctx, query, args...)
}
