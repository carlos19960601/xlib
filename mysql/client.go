package mysql

import (
	"database/sql"

)

type Client struct {
	*sql.DB
	options Options
}

func NewMySQL(options ...Option) *Client {
	opts := NewOptions(options...)

	mysql := &Client{
		options: opts,
	}

	db, err := sql.Open("mysql", mysql.options.DSN)
	if err != nil {
		panic("无法链接数据库, dburl: " + mysql.options.DSN)
	}

	db.SetMaxOpenConns(mysql.options.MaxOpenConns)
	db.SetMaxIdleConns(mysql.options.MaxIdleConns)
	db.SetConnMaxLifetime(mysql.options.ConnMaxLifetime)

	mysql.DB = db
	return mysql
}
