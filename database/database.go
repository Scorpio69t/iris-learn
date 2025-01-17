package database

import "iris-learn/environment"

type DB interface {
	Init(source string) error
}

// NewDB returns a new DB based on the environment.
func NewDB(env environment.Env) DB {
	switch env {
	case environment.PROD:
		return &mysql{DB: nil}
	case environment.DEV:
		return &sqlite{DB: nil}
	default:
		panic("Unknown environment")
	}
}
