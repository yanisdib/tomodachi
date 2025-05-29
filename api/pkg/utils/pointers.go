package utils

import (
	"database/sql"
)

func PtrIfValid[T any](valid bool, val T) *T {
	if valid {
		return &val
	}

	return nil
}

func NullStringPtr(ns sql.NullString) *string {
	return PtrIfValid(ns.Valid, ns.String)
}

func NullInt64Ptr(ni sql.NullInt64) *int64 {
	return PtrIfValid(ni.Valid, ni.Int64)
}

func NullInt16Ptr(ni sql.NullInt16) *int16 {
	return PtrIfValid(ni.Valid, ni.Int16)
}

func NullBoolPtr(nb sql.NullBool) *bool {
	return PtrIfValid(nb.Valid, nb.Bool)
}
