package serializers

import "github.com/google/uuid"

type User struct {
	Id             uuid.UUID `json:"id"`
	PassportNumber string    `json:"passportNumber"`
	PassportSeries string    `json:"passportSeries"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Patronymic     string    `json:"patronymic"`
	Address        string    `json:"address"`
}

type CreateUserRequest struct {
	Passport Passport `json:"passport" binding:"required"`
}

type CreateUserResponse struct {
	Id uuid.UUID `json:"id"`
}

type GetUsersRequest struct {
	Id             *uuid.UUID      `form:"id"`
	PassportNumber *PassportNumber `form:"passportNumber"`
	PassportSeries *PassportSeries `form:"passportSeries"`
	Name           *string         `form:"name"`
	Surname        *string         `form:"surname"`
	Patronymic     *string         `form:"patronymic"`
	Address        *string         `form:"address"`
	Page           *int            `form:"page"`
	PageSize       *int            `form:"pageSize"`
}

type GetUsersResponse struct {
	Info []User `json:"info"`
}

type UpdateUserRequest struct {
	PassportNumber *PassportNumber `json:"passportNumber"`
	PassportSeries *PassportSeries `json:"passportSeries"`
	Name           *string         `json:"name"`
	Surname        *string         `json:"surname"`
	Patronymic     *string         `json:"patronymic"`
	Address        *string         `json:"address"`
}

type UpdateUserResponse struct {
	Status string `json:"status"`
}

type DeleteUserRequest struct {
	Id uuid.UUID `json:"id"`
}
type DeleteUserResponse struct {
	Status string `json:"status"`
}
