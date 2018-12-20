package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Car struct {
	ID     int    `json:"id"`
	Vin    string `json:"vin,omitempty"`
	Model  string `json:"model,omitempty"`
	Maker  string `json:"maker,omitempty"`
	Year   int    `json:"year,string"`
	Msrp   int    `json:"msrp,omitempty,string"`
	Status string `json:"status,omitempty"`
	Booked bool   `json:"booked"`
	Listed bool   `json:"listed"`
}

var cars []Car

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if i, err := strconv.Atoi(params["id"]); err == nil {
		for _, item := range cars {
			if item.ID == i {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}

	json.NewEncoder(w).Encode(cars)
}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(cars)
}

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var car Car
	decoder.Decode(&car)
	car.ID = len(cars) + 1
	cars = append(cars, car)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(cars)

}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range cars {
		if i, err := strconv.Atoi(params["id"]); err == nil {
			if item.ID == i {
				cars = append(cars[:index], cars[index+1:]...)
				break
			}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			json.NewEncoder(w).Encode(cars)
		}

	}
}

func main() {
	router := mux.NewRouter()
	cars = append(cars, Car{1, "MNBUMF050FW496402", "320i", "BMW", 2013,
		10000, "Ordered", true, true})
	cars = append(cars, Car{2, "4JDBLMF080FW468775", "Carmry", "Toyota", 2015,
		12000, "In stock", true, false})
	cars = append(cars, Car{3, "TFBAXXMAWAFS71274", "Focus", "Ford", 2016,
		13000, "Ordered", false, true})
	cars = append(cars, Car{4, "G3SBUMF080FW470449", "Civic", "Honda", 2016,
		14000, "Sold", false, false})
	router.HandleFunc("/cars", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/cars/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/cars", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/cars/{id}", DeletePersonEndpoint).Methods("DELETE")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(corsObj, headersOk, methodsOk)(router)))
}
