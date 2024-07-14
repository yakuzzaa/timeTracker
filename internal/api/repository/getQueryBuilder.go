package repository

import (
	"errors"
	"timeTracker/internal/api/serializers"

	"github.com/google/uuid"
)

type UserQueryBuilder interface {
	WithID(id *uuid.UUID) UserQueryBuilder
	WithPassportNumber(passportNumber *serializers.PassportNumber) (UserQueryBuilder, error)
	WithPassportSeries(passportSeries *serializers.PassportSeries) (UserQueryBuilder, error)
	WithName(name *string) UserQueryBuilder
	WithSurname(surname *string) UserQueryBuilder
	WithPatronymic(patronymic *string) UserQueryBuilder
	WithAddress(address *string) UserQueryBuilder
	WithPagination(page, pageSize *int) (UserQueryBuilder, error)
	Build() (string, []interface{})
}

type userQueryBuilder struct {
	conditions []string
	args       []interface{}
	page       int
	pageSize   int
}

func NewUserQueryBuilder() UserQueryBuilder {
	return &userQueryBuilder{}
}

func (qb *userQueryBuilder) WithID(id *uuid.UUID) UserQueryBuilder {
	if id != nil {
		qb.conditions = append(qb.conditions, "id = ?")
		qb.args = append(qb.args, id)
	}
	return qb
}

func (qb *userQueryBuilder) WithPassportNumber(passportNumber *serializers.PassportNumber) (UserQueryBuilder, error) {
	if passportNumber != nil {
		if err := passportNumber.Validate(); err != nil {
			return qb, err
		}
		qb.conditions = append(qb.conditions, "passport_number = ?")
		qb.args = append(qb.args, *passportNumber)
	}
	return qb, nil
}

func (qb *userQueryBuilder) WithPassportSeries(passportSeries *serializers.PassportSeries) (UserQueryBuilder, error) {
	if passportSeries != nil {
		if err := passportSeries.Validate(); err != nil {
			return qb, err
		}
		qb.conditions = append(qb.conditions, "passport_series = ?")
		qb.args = append(qb.args, passportSeries)
	}
	return qb, nil
}

func (qb *userQueryBuilder) WithName(name *string) UserQueryBuilder {
	if name != nil {
		qb.conditions = append(qb.conditions, "name = ?")
		qb.args = append(qb.args, name)
	}
	return qb
}

func (qb *userQueryBuilder) WithSurname(surname *string) UserQueryBuilder {
	if surname != nil {
		qb.conditions = append(qb.conditions, "surname = ?")
		qb.args = append(qb.args, surname)
	}
	return qb
}

func (qb *userQueryBuilder) WithPatronymic(patronymic *string) UserQueryBuilder {
	if patronymic != nil {
		qb.conditions = append(qb.conditions, "patronymic = ?")
		qb.args = append(qb.args, patronymic)
	}
	return qb
}

func (qb *userQueryBuilder) WithAddress(address *string) UserQueryBuilder {
	if address != nil {
		qb.conditions = append(qb.conditions, "address = ?")
		qb.args = append(qb.args, address)
	}
	return qb
}

func (qb *userQueryBuilder) WithPagination(page, pageSize *int) (UserQueryBuilder, error) {
	if page != nil && pageSize != nil {
		qb.page = *page
		qb.pageSize = *pageSize
	}

	if page != nil && pageSize == nil {
		return nil, errors.New("pageSize must be specified")
	}

	if page == nil && pageSize != nil {
		return nil, errors.New("page must be specified")
	}
	return qb, nil
}

func (qb *userQueryBuilder) Build() (string, []interface{}) {
	sql := "SELECT * FROM users WHERE 1=1"

	if len(qb.conditions) > 0 {
		for _, cond := range qb.conditions {
			sql += " AND " + cond
		}
	}

	if qb.page != 0 && qb.pageSize != 0 {
		sql += " OFFSET ? LIMIT ?"
		qb.args = append(qb.args, (qb.page-1)*qb.pageSize, qb.pageSize)
	}

	return sql, qb.args
}
