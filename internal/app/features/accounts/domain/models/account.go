package models

type Account struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	UserID      int     `json:"user_id"`
	Summary     float64 `json:"summary"`
	Description string  `json:"description,omitempty"`
}
