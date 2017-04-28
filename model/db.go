package model

type db interface {
	People() ([]*Person, error)
}

type Person struct {
	First, Last string
}
