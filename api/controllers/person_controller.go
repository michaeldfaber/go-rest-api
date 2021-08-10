package controllers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"

	"go-rest-api/api/models"
	"go-rest-api/api/responses"
)

func (server *Server) CreatePerson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

    var person *models.Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	person.Prepare()

	personCreated, err := person.SavePerson(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	responses.JSON(w, http.StatusCreated, personCreated)
}

func (server *Server) GetAllPersons(w http.ResponseWriter, r *http.Request) {
    var person *models.Person

	persons, err := person.FindAllPersons(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, persons)
}

// func (server *Server) GetPerson(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	idString := vars["id"]
// 	id, _ := strconv.Atoi(idString)
// 	var notFound = true
// 	for _, person := range Persons {
// 		if person.Id == id {
// 			notFound = false
// 			json.NewEncoder(w).Encode(person)
// 		}
// 	}
// 	if notFound {
// 		fmt.Fprintf(w, "No person with that Id was found.")
// 	}
// }

// func (server *Server) UpdatePerson(w http.ResponseWriter, r *http.Request) {
// 	requestBody, _ := ioutil.ReadAll(r.Body)
// 	var updatedPerson *models.Person 
// 	json.Unmarshal(requestBody, &updatedPerson)
// 	personIndex := -1
//     for i, person := range Persons {
//         if person.Id == updatedPerson.Id {
// 			personIndex = i
//         }
// 	}
// 	if personIndex != -1 {
// 		Persons[personIndex] = updatedPerson
// 	} else {
// 		fmt.Fprintf(w, "No person with that Id was found. No changes were made.")
// 	}
// }

// func (server *Server) DeletePerson(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     idString := vars["id"]
// 	id, _ := strconv.Atoi(idString)
//     for i, person := range Persons {
//         if person.Id == id {
//             Persons = append(Persons[:i], Persons[i+1:]...)
//         }
//     }
// }

// func (server *Server) CreateRandomPerson(w http.ResponseWriter, r *http.Request) {
// 	randomPerson := getRandomPerson()
// 	var person *models.Person
// 	person.Id = Persons[len(Persons)-1].Id + 1
// 	person.Gender = randomPerson.Gender
// 	person.FirstName = randomPerson.Name.First
// 	person.LastName = randomPerson.Name.Last
// 	person.Age = randomPerson.Dob.Age
// 	Persons = append(Persons, person)
//     json.NewEncoder(w).Encode(person)
// }

// func getRandomPerson() models.RandomPerson {
// 	randomPersonHttpResponse, _ := http.Get("https://randomuser.me/api")
// 	randomPersonHttpResponseBody, _ := ioutil.ReadAll(randomPersonHttpResponse.Body)
//     var randomPersonResponse *models.RandomPersonResponse
// 	json.Unmarshal(randomPersonHttpResponseBody, &randomPersonResponse)
// 	return randomPersonResponse.Results[0]
// }