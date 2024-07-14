package models

import "github.com/google/uuid"

type User struct {
	Id             uuid.UUID
	PassportNumber string
	PassportSeries string
	Name           string
	Surname        string
	Patronymic     string
	Address        string
}
