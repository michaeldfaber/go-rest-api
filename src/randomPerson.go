package gorestapi

type Name struct {
	First string `json:"first"`
	Last string `json"last"`
}

type Dob struct {
	Age int `json:"age"`
}

type RandomPerson struct {
	Gender string `json:"gender"`
	Name Name `json:"name"`
	Dob Dob `json:"dob"`
}

type RandomPersonResponse struct {
	Results []RandomPerson `json:"results"`
}