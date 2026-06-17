package helper

import (
	"database/sql"
	"time"

	"github.com/sqlc-dev/pqtype"
)

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	nullStr := sql.NullString{
		String: s,
		Valid:  true,
	}
	return nullStr
}

func ToRawMessage(s string) pqtype.NullRawMessage {
	if s == "" {
		return pqtype.NullRawMessage{Valid: false}
	}
	val := pqtype.NullRawMessage{
		RawMessage: []byte(s),
		Valid:      true,
	}
	return val
}

func StringToNullTime(s string) sql.NullTime {
	layout := "2006-01-02 15:04:05"
	if s == "" {
		return sql.NullTime{Valid: false}
	}

	t, err := time.Parse(layout, s)
	if err != nil {
		return sql.NullTime{Valid: false}
	}

	return sql.NullTime{Time: t, Valid: true}
}
