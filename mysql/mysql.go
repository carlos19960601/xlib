package mysql

import (
	"database/sql"
)

type MySQL struct {
	db      *sql.DB
	options Options
}

func NewMySQL(options ...Option) *MySQL {
	opts := NewOptions(options...)

	mysql := &MySQL{
		options: opts,
	}

	db, err := sql.Open("mysql", mysql.options.DSN)
	if err != nil {
		panic("无法链接数据库, dburl: " + mysql.options.DSN)
	}

	db.SetMaxOpenConns(mysql.options.MaxOpenConns)
	db.SetMaxIdleConns(mysql.options.MaxIdleConns)
	db.SetConnMaxLifetime(mysql.options.ConnMaxLifetime)

	mysql.db = db
	return mysql
}
