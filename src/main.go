package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	// "strconv"
	"io/ioutil"
)

type Person struct {
	Id string `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Age int `json:"Age"`
}
var Persons []Person

func defaultPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func returnAllPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Persons)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, _ := ioutil.ReadAll(r.Body)
    var person Person 
    json.Unmarshal(requestBody, &person)
    Persons = append(Persons, person)
    json.NewEncoder(w).Encode(person)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Tried to have Id as an int but had issues converting it from string in URL

	// id, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	for _, person := range Persons {
	// 		if person.Id == id {
	// 			json.NewEncoder(w).Encode(person)
	// 		}
	// 	}
	// } else {
	// 	fmt.Fprintf(w, "No person with that Id was found.")
	// }

	// .Where() / .Find() / .filter() instead?
	var notFound = true
	for _, person := range Persons {
		if person.Id == id {
			notFound = false
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(person)
		}
	}

	if notFound {
		fmt.Fprintf(w, "No person with that Id was found.")
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id := vars["id"]
    for i, person := range Persons {
        if person.Id == id {
            Persons = append(Persons[:i], Persons[i+1:]...)
        }
    }
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	
	myRouter.HandleFunc("/", defaultPage)
	myRouter.HandleFunc("/all", returnAllPersons)
	myRouter.HandleFunc("/create", createPerson).Methods("POST")
	myRouter.HandleFunc("/person/{id}", getPerson)
	myRouter.HandleFunc("/delete/{id}", deletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	// Initial data, no time for a real db for now
	Persons = []Person {
		Person { Id: "1", FirstName: "Michael", LastName: "Faber", Age: 24 },
		Person { Id: "2", FirstName: "Bob", LastName: "Johnson", Age: 30 },
		Person { Id: "3", FirstName: "John", LastName: "Smith", Age: 45 },
	}
    handleRequests()
}