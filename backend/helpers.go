package main

import (
	"errors"

	sqlite "modernc.org/sqlite"
	sqlLib "modernc.org/sqlite/lib"
)

func IsUniqueConstraintError(err error) bool {
	var sqliteErr *sqlite.Error
	if errors.As(err, &sqliteErr) {
		return sqliteErr.Code() == sqlLib.SQLITE_CONSTRAINT_UNIQUE
	}
	return false
}
