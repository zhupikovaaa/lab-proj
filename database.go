package main
import (
	"database/sql"
	"fmt"
)
type Car struct{
	Mark string
	Country string
	Price int
	Year int
}
const (
	DB_USER = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME = "lab"
)
func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS car (mark_name text, country text,price integer, release_year integer);"); err != nil {
	return err
}
return nil
}
func dbAddCar(mark string, country string,price int, year int) error {
	sqlstmt := "INSERT INTO car VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(sqlstmt, mark, country, price, year)
	if err != nil {
		return err
	}
	return nil
}
func dbGetCars() ([]Car, error) {
	var car []Car
	stmt, err := db.Prepare("SELECT mark_name, country, price, release_year FROM car")
	if err != nil {
		return car, err
	}
	res, err := stmt.Query()
	if err != nil {
		return car, err
	}
	var tempCar Car
	for res.Next() {
		err = res.Scan(&tempCar.Mark, &tempCar.Country, &tempCar.Price, &tempCar.Year)
		if err != nil {
			return car, err
		}
		car = append(car, tempCar)
	}
	return car, err
}

