package db

type DB interface {
	Close()
	Info()
}

