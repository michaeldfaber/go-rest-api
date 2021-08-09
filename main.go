package gorestapi

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strconv"
	
	"github.com/michaeldfaber/go-rest-api/api/models"
)

var Persons []Person

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
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)
	var notFound = true
	for _, person := range Persons {
		if person.Id == id {
			notFound = false
			json.NewEncoder(w).Encode(person)
		}
	}
	if notFound {
		fmt.Fprintf(w, "No person with that Id was found.")
	}
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var updatedPerson Person 
	json.Unmarshal(requestBody, &updatedPerson)
	personIndex := -1
    for i, person := range Persons {
        if person.Id == updatedPerson.Id {
			personIndex = i
        }
	}
	if personIndex != -1 {
		Persons[personIndex] = updatedPerson
	} else {
		fmt.Fprintf(w, "No person with that Id was found. No changes were made.")
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    idString := vars["id"]
	id, _ := strconv.Atoi(idString)
    for i, person := range Persons {
        if person.Id == id {
            Persons = append(Persons[:i], Persons[i+1:]...)
        }
    }
}

func createRandomPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	randomPerson := getRandomPerson()
	var person Person
	person.Id = Persons[len(Persons)-1].Id + 1
	person.Gender = randomPerson.Gender
	person.FirstName = randomPerson.Name.First
	person.LastName = randomPerson.Name.Last
	person.Age = randomPerson.Dob.Age
	Persons = append(Persons, person)
    json.NewEncoder(w).Encode(person)
}

func getRandomPerson() RandomPerson {
	randomPersonHttpResponse, _ := http.Get("https://randomuser.me/api")
	randomPersonHttpResponseBody, _ := ioutil.ReadAll(randomPersonHttpResponse.Body)
    var randomPersonResponse RandomPersonResponse
	json.Unmarshal(randomPersonHttpResponseBody, &randomPersonResponse)
	return randomPersonResponse.Results[0]
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", returnAllPersons)
	myRouter.HandleFunc("/create", createPerson).Methods("POST")
	myRouter.HandleFunc("/create/random", createRandomPerson).Methods("POST")
	myRouter.HandleFunc("/person/{id}", getPerson)
	myRouter.HandleFunc("/update", updatePerson).Methods("PUT")
	myRouter.HandleFunc("/delete/{id}", deletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Persons = []Person {
		Person { Id: 1, Gender: "male", FirstName: "Michael", LastName: "Faber", Age: 24 },
		Person { Id: 2, Gender: "male", FirstName: "Bob", LastName: "Johnson", Age: 30 },
		Person { Id: 3, Gender: "female", FirstName: "Jane", LastName: "Smith", Age: 45 },
	}
    handleRequests()
}