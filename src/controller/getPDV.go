package controller

import (
	"fmt"
	db "lao-pdv/src/config"
	"lao-pdv/src/model"
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
)

func GetProvince(c echo.Context) error {
	db := db.Conn()
	// create the select sql query
	sqlStatement := sq.Select("*").From("provinces")

	// close database
	defer db.Close()

	// execute the sql statement
	rows, err := sqlStatement.RunWith(db).Query()

	// An Province slice to hold data from returned rows.
	var provinces []model.Province

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var province model.Province

		err = rows.Scan(&province.ID, &province.Name, &province.Name_en)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the province in the provinces slice
		provinces = append(provinces, province)
	}
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, provinces)
}


func GetDistrict(c echo.Context) error {
	db := db.Conn()
	sqlStatement := sq.Select("*").From("districts")
	//sqlStatement := sq.Select("districts.id, districts.name, districts.name_en, provinces.name").From("districts").Join("provinces on districts.province_id = provinces.id")

	// close database
	defer db.Close()

	rows, err := sqlStatement.RunWith(db).Query()

	var districts []model.District

	for rows.Next() {
		var district model.District
		err = rows.Scan(&district.ID, &district.Name, &district.Name_en, &district.Province_id)
		
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		districts = append(districts, district)
	}
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, districts)
}

func GetVillage(c echo.Context) error {
	db := db.Conn()

	sqlStatement := sq.Select("*").From("villages")

	// close database
	defer db.Close()

	rows, err := sqlStatement.RunWith(db).Query()

	var villages []model.Village

	for rows.Next() {
		var village model.Village
		err = rows.Scan(&village.ID, &village.Name, &village.Name_en, &village.District_id)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		villages = append(villages, village)
	}
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, villages)
}