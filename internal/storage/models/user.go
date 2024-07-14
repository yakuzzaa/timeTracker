package models

import "github.com/google/uuid"

type User struct {
	Id             uuid.UUID `json:"id" db:"id"`
	PassportNumber string    `json:"passport_number" db:"passport_number"`
	PassportSeries string    `json:"passport_series" db:"passport_series"`
	Name           string    `json:"name" db:"name"`
	Surname        string    `json:"surname" db:"surname"`
	Patronymic     string    `json:"patronymic" db:"patronymic"`
	Address        string    `json:"address" db:"address"`
}
