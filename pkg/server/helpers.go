package server

import (
	"fmt"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func formatAsRfc3339String(t time.Time) string {
	return t.Format(time.RFC3339)
}
