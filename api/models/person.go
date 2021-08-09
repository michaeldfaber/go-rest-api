package models

type Person struct {
	Id int `json:"Id"`
	Gender string `json:"Gender"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Age int `json:"Age"`
}