package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Planet is a struct for the plane
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

// Person is a struct for the person
type Person struct {
	Name         string `json:"name"`
	HomeworldUrl string `json:"homeworld"`
	Homeworld    Planet
}

// AllPeople is a struct for the people
type AllPeople struct {
	People []Person `json:"results"`
}

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldUrl)

	if err != nil {
		log.Println("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = io.ReadAll(res.Body); err != nil {
		log.Println("Error reading response body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPeople")
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print("Failed to reqest start wars people")
	}

	bytes, err := io.ReadAll(res.Body)
	fmt.Println(string(bytes))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print("Failed to read start wars people")
	}

	var people AllPeople

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	for _, pers := range people.People {
		pers.getHomeworld()
		fmt.Println(pers)
	}
}

func main() {
	http.HandleFunc("/", getPeople)

	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
