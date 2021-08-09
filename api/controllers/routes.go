package controllers

import "go-rest-api/api/middleware"

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Person Routes
	s.Router.HandleFunc("/person/all", middlewares.SetMiddlewareJSON(s.GetAllPersons))
	s.Router.HandleFunc("/person/{id}", middlewares.SetMiddlewareJSON(s.GetPerson))
	s.Router.HandleFunc("/person/create", middlewares.SetMiddlewareJSON(s.CreatePerson)).Methods("POST")
	s.Router.HandleFunc("/person/create/random", middlewares.SetMiddlewareJSON(s.CreateRandomPerson)).Methods("POST")
	s.Router.HandleFunc("/person/update", middlewares.SetMiddlewareJSON(s.UpdatePerson)).Methods("PUT")
	s.Router.HandleFunc("/person/delete/{id}", middlewares.SetMiddlewareJSON(s.DeletePerson)).Methods("DELETE")
}