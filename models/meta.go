package models

type Meta struct {
	Count int   `json:"count"`
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}
