package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kwamekyeimonies/stocks-api/database"
	"github.com/kwamekyeimonies/stocks-api/models"
)

type response struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.Create_connection()
	defer db.Close()
	stock := &models.Stock{
		StockID: uuid.New().String(),
	}

	_ = json.NewDecoder(r.Body).Decode(&stock)

	_, err := db.Model(stock).Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// res := response{
	// 	ID: stock.StockID,
	// 	Message: "stock created Successfully.....",
	// }

	json.NewEncoder(w).Encode(stock)

}

func GetAllStock(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.Create_connection()
	defer db.Close()

	var stock []models.Stock
	if err := db.Model(&stock).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Create_connection()
	defer db.Close()
	params := mux.Vars(r)

	stockId := params["id"]
	stock := &models.Stock{StockID: stockId}

	if err := db.Model(stock).WherePK().Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Create_connection()
	defer db.Close()
	params := mux.Vars(r)

	stockId := params["id"]
	stock := &models.Stock{StockID: stockId}

	result, err := db.Model(stock).Where("id = ?", stockId).Delete()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.Create_connection()
	defer db.Close()
	params := mux.Vars(r)
	stockId := params["id"]

	stock := &models.Stock{StockID: stockId}
	_ = json.NewDecoder(r.Body).Decode(&stock)
	_, err := db.Model(stock).WherePK().Set("name= ?, price= ?,company= ?", stock.Name, stock.Price, stock.Company).Update()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(stock)
}
