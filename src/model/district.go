package model

type District struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Name_en string `json:"name_en"`
	Province_id int `json:"province_id"`
	//Province_id string `json:"province_id"`
}