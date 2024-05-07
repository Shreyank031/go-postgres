package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Shreyank031/go-postgres/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the postgres driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) { //post

	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body: %v", err)
	}
	insertId := insertStock(stock)

	res := response{
		ID:      insertId,
		Message: "Stock created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int: %v", err)
	}
	res, err := getStocksById(int64(id))
	if err != nil {
		log.Fatalf("Unable to fetch stocks from id: %v", err)
	}

	json.NewEncoder(w).Encode(res)

}

func GetAllStocks(w http.ResponseWriter, h *http.Request) {
	res, err := getAllStock()
	if err != nil {
		log.Fatalf("Unable to get all stocks from database: %v", err)
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert id from string to int: %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode from json to struct: %v", err)
	}

	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock updated successfully. Total rows affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStockById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int: %v", err)
	}
	deletedRows, err := deleteStock(int64(id))
	if err != nil {
		log.Fatal(err)
	}
	msg := fmt.Sprintf("Stock deleted successfully. Total record/rows affected: %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatements := `INSERT INTO stockdb.stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stocksid`

	var id int64

	err := db.QueryRow(sqlStatements, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	fmt.Println()
	return id
}

func getStocksById(id int64) (models.Stock, error) {

	db := createConnection()
	defer db.Close()

	var stock models.Stock
	sqlStatement := `SELECT * FROM stockdb.stocks WHERE stocksid=$1`
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}
	return stock, err
}

func getAllStock() ([]models.Stock, error) {

	db := createConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM stockdb.stocks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the Query: %v", err)
	}

	defer rows.Close()
	var stocks []models.Stock

	for rows.Next() {
		var stock models.Stock

		err := rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the rows: %v", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, err

}

func updateStock(id int64, stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE stockdb.stocks SET name=$2, price=$3, company=$4 WHERE stocksid=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to execute the Query: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the effected rows: %v", err)
	}
	return rowsAffected

}
func deleteStock(id int64) (int64, error) {

	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stockdb.stocks WHERE stocksid=$1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return 0, fmt.Errorf("error deleting stock: %w", err)
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking affected rows: %v", err)
	}
	return affectedRows, err
}
