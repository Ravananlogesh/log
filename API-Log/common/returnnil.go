package common

import (
	"database/sql"
	"strings"
)



func ReturnNil(value string) interface{} {
	if value == "" {
		return nil
	} else {
		return value
	}
}

func NewNullString(s string) sql.NullString {
	if strings.EqualFold(s, "null") {
		return sql.NullString{}
	}
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
