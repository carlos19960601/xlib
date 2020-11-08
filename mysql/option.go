package mysql

import (
	"time"
)

var (
	DefaultMaxOpenConns = 10
)

type Options struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type Option func(*Options)

func MaxOpenConns(a int) Option {
	return func(o *Options) {
		o.MaxOpenConns = a
	}
}

func MaxIdleConns(a int) Option {
	return func(o *Options) {
		o.MaxIdleConns = a
	}
}

func ConnMaxLifetime(t time.Duration) Option {
	return func(o *Options) {
		o.ConnMaxLifetime = t
	}
}

func DSN(a string) Option {
	return func(o *Options) {
		o.DSN = a
	}
}

func NewOptions(options ...Option) Options {
	opts := Options{
		MaxOpenConns: DefaultMaxOpenConns,
	}
	for _, o := range options {
		o(&opts)
	}
	return opts
}
