package mongoDB

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func DbConnection() {
	url := `mongodb+srv://` + os.Getenv("DB_USERNAME") + `:` + os.Getenv("DB_PASSWORD") + `@stocks.xhvknoe.mongodb.net/?retryWrites=true&w=majority`
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(os.Getenv("DB_NAME"))
	DB = db

}
