package model

import "database/sql"

type Village struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Name_en     sql.NullString `json:"name_en"`
	District_id int            `json:"district_id"`
}