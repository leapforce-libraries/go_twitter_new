package models

import (
	"strconv"
	"time"
)

const (
	DateLayout string = "2006-01-02T15:04:05.999Z"
)

func parseInt64(s *string) (*int64, error) {
	if s == nil {
		return nil, nil
	}

	id, err := strconv.ParseInt(*s, 10, 64)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func parseTime(s string) (*time.Time, error) {
	t, err := time.Parse(DateLayout, s)

	if err != nil {
		return nil, err
	}

	return &t, err
}
