package model

type db interface {
	SelectPeople() ([]*Person, error)
}
