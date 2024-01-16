package service

import "time"

type Servise struct {
}

func New() *Servise {
	return &Servise{}
}

func (s *Servise) DaysLeft() int64 {
	date := time.Date(2025, time.June, 1, 0, 0, 0, 0, time.UTC)

	duration := time.Until(date)

	return int64(duration.Hours()) / 24
}
