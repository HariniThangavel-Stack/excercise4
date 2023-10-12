package controller

import (
	"context"
	"encoding/json"
	"excercise4/dbiface"
	"excercise4/model"
	"excercise4/mongoDB"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Base(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Listening on port 3000"))
}

func GetData(collection dbiface.CollectionAPI, data string) (*mongo.Cursor, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{"symbol": data})
	if err != nil {
		log.Fatal(err)
	}
	return cursor, nil
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var results []model.StockMarketData
	marketDataCollection := mongoDB.DB.Collection("marketData")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := marketDataCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	json.NewEncoder(w).Encode(results)
}

func GetFileredStocks(w http.ResponseWriter, r *http.Request) {
	symbol := r.FormValue("symbol")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var results []model.StockMarketData
	marketDataCollection := mongoDB.DB.Collection("marketData")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := GetData(marketDataCollection, symbol)
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(results)
	defer cursor.Close(ctx)
}

func InsertData(collection dbiface.CollectionAPI, data model.StockMarketData) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return cursor, nil
}

func InsertStock(w http.ResponseWriter, r *http.Request) {
	var stockData model.StockMarketData
	err := json.NewDecoder(r.Body).Decode(&stockData)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	marketDataCollection := mongoDB.DB.Collection("marketData")
	cursor, _ := InsertData(marketDataCollection, stockData)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(cursor.InsertedID)
}

func UpdateData(collection dbiface.CollectionAPI, find string, body model.UpdateBody) *mongo.SingleResult {
	filter := bson.D{{"symbol", find}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"symbol", body.Symbol}}}}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor := collection.FindOneAndUpdate(ctx, filter, update, &returnOpt)
	return cursor
}

func FindAndUpdateStockName(w http.ResponseWriter, r *http.Request) {
	var body model.UpdateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}
	symbol := r.FormValue("symbol")
	var result model.StockMarketData

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	marketDataCollection := mongoDB.DB.Collection("marketData")
	cursor := UpdateData(marketDataCollection, symbol, body)
	cursor.Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func ReplaceOneDoc(collection dbiface.CollectionAPI, find string, body model.StockMarketData) (*mongo.UpdateResult, error) {
	filter := bson.D{{"symbol", find}}
	update := bson.M{
		"symbol":                body.Symbol,
		"open":                  body.Open,
		"high":                  body.High,
		"low":                   body.Low,
		"previousclose":         body.PreviousClose,
		"ltp":                   body.Ltp,
		"change":                body.Change,
		"percentagechange":      body.PercentageChange,
		"volume":                body.Volume,
		"value":                 body.Value,
		"yearhigh":              body.YearHigh,
		"yearlow":               body.YearLow,
		"yearpercentagechange":  body.YearPercentageChange,
		"monthpercentagechange": body.MonthPercentageChange,
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.ReplaceOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return cursor, nil
}

func FindAndUpdateStock(w http.ResponseWriter, r *http.Request) {
	var body model.StockMarketData

	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}
	symbol := r.FormValue("symbol")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	marketDataCollection := mongoDB.DB.Collection("marketData")
	if r.Method == "PUT" || r.Method == "put" {
		cursor, e := ReplaceOneDoc(marketDataCollection, symbol, body)
		if e != nil {
			fmt.Print(e)
		}
		json.NewEncoder(w).Encode(cursor.ModifiedCount)
	}
}

func DeleteData(collection dbiface.CollectionAPI, data string) (*mongo.DeleteResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.DeleteOne(ctx, bson.M{"symbol": data})
	if err != nil {
		log.Fatal(err)
	}
	return cursor, nil
}

func FindAndDeleteStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	symbol := r.FormValue("symbol")
	marketDataCollection := mongoDB.DB.Collection("marketData")
	cursor, _ := DeleteData(marketDataCollection, symbol)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(cursor.DeletedCount)
}
