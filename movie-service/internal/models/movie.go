package models

import (
	"fmt"
	"time"
)

type Movie struct {
	ID       int
	Title    string
	Description string
	Duration time.Duration
	Age_Limit int
	CreatedAt     time.Time
}

func (m *Movie) Validate() error {
	if len(m.Title) == 0 && len(m.Description) == 0 && m.Duration == 0 && m.Age_Limit == 0 {
		return fmt.Errorf("Validation error: title, description, duration and age_limit are required")
	}
	return nil
}

type UpdateMovie struct {
	ID          int
	Title       string
	Description string
	Duration time.Duration
	Age_Limit int
}

func (up *UpdateMovie) Validate() error {
	if len(up.Title) == 0 && len(up.Description) == 0 && up.Duration == 0 && up.Age_Limit == 0 {
		return fmt.Errorf("Validation error: title, description, duration and age_limit are required")
	}

	return nil
}
