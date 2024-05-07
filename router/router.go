package router

import (
	"github.com/Shreyank031/go-postgres/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/stock/{id}", middleware.GetStockById).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/deletestock/{id}", middleware.DeleteStockById).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/newstock/{id}", middleware.CreateStock).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")

	return r
}
