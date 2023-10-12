package main

import (
	"context"
	"encoding/json"
	"excercise4/controller"
	"excercise4/model1"
	"excercise4/mongoDB"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	parser "github.com/HariniThangavel-Stack/market-data-parser-golang"
	"github.com/HariniThangavel-Stack/market-data-parser-golang/model"
	"github.com/gorilla/mux"
)

func init() {
	mongoDB.DbConnection()
}

func main() {
	if _, err := os.Stat("./examples/MW-NIFTY-BANK-05-Aug-2021.csv1"); err == nil {
		options := model.Options{OutputFormat: "json", FilePath: "./examples/MW-NIFTY-BANK-05-Aug-2021.csv"}
		marketData := parser.MarketDataParser(options)
		var mData []interface{}
		for _, i := range marketData {
			var res model1.StockMarketData
			bytes := []byte(i)
			json.Unmarshal(bytes, &res)
			mData = append(mData, (res))
		}
		marketDataCollection := mongoDB.DB.Collection("marketData")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		result, err := marketDataCollection.InsertMany(ctx, mData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.InsertedIDs)
	}

	route := mux.NewRouter()
	server := route.PathPrefix("/api").Subrouter()

	server.HandleFunc("/test", controller.Base).Methods("GET")
	server.HandleFunc("/getStocks", controller.GetAllStocks).Methods("GET", "OPTIONS")
	server.HandleFunc("/getStocksBySymbol", controller.GetFileredStocks).Methods("GET")
	server.HandleFunc("/insertStock", controller.InsertStock).Methods("POST", "OPTIONS")
	server.HandleFunc("/findAndUpdateStock", controller.FindAndUpdateStockName).Methods("PATCH")
	server.HandleFunc("/findAndUpdateStock", controller.FindAndUpdateStock).Methods("PUT", "OPTIONS")
	server.HandleFunc("/deleteStock", controller.FindAndDeleteStock).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8000", server))

}
