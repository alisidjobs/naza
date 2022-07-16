package router

import (
	"dnd/server/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/perso", middleware.GetAllPerso).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/perso", middleware.CreatePerso).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/modifPerso/{id}", middleware.ModifyPerso).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletePerso/{id}", middleware.DeletePerso).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllPerso", middleware.DeleteAllPerso).Methods("DELETE", "OPTIONS")
	return router
}
