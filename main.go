package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	_"github.com/lib/pq"

	"database/sql"
	"fmt"
	"os"
)
var db *sql.DB
func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		cars, err := dbGetCars()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, cars)
	}
}
func addCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		mark := r.Form.Get("mark")
		country:= r.Form.Get("country")
		year, errI := strconv.Atoi(r.Form.Get("year"))
		price, errI := strconv.Atoi(r.Form.Get("price"))
		err := dbAddCar(mark, country, year, price)
		if err != nil || errI != nil {
			log.Fatal(err)
		}
	}
}
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println(port)
	}
	return ":" + port
}
func main() {
	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addCarHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}