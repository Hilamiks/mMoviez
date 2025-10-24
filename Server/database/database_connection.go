package database

import (
	"fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"log"
	"os"
)

func DBInstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URI")

	if MongoDb == "" {
		log.Fatal("MONGODB_URI not set in environment")
	}

	fmt.Println("MongoDB URI: ", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	fmt.Println("DB Name:", dbName)

	return Client.Database(dbName).Collection(collectionName)
}
