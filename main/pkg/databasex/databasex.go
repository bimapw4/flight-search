package databasex

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrCanceled            = "pq: canceling statement due to user request"
	ErrUniqueViolation     = Error("unique_violation")
	ErrNullValueNotAllowed = Error("null_value_not_allowed")
	ErrorUndefinedTable    = Error("undefined_table")
	ErrNoRowsFound         = Error("sql: no rows in result set")
)

func ParsePostgreSQLError(err error) error {
	// Parse by type
	if err == sql.ErrNoRows {
		return ErrNoRowsFound
	}
	switch et := err.(type) {
	case *pq.Error:
		switch et.Code {
		case "23505":
			return ErrUniqueViolation
		case "42P01":
			return ErrorUndefinedTable
		case "22004":
			return ErrNullValueNotAllowed
		}
	}
	// Parse by message
	switch err.Error() {
	case ErrCanceled:
		return context.Canceled
	}
	return err
}
