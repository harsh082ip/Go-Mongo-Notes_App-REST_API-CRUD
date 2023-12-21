package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// "encoding/json"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// start-restaurant-struct
type Notes struct {
	Id    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title,omitempty"`
	Desc  string             `bson:"desc,omitempty"`
	Time  string             `bson:"time,omitempty"`
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)
	fmt.Println(id["id"])
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin findOne
	coll := client.Database("sample_restaurants").Collection("restaurants")

	// Creates a query filter to match documents in which the "name" is
	// "Bagels N Buns"
	objectID, err := primitive.ObjectIDFromHex(id["id"])
	if err != nil {
		panic(err)
	}
	fmt.Println(objectID)
	// filter := bson.D{{"title", "golang"}}
	filter := bson.D{{"_id", objectID}}

	// Retrieves the first matching document
	var result Notes
	err = coll.FindOne(context.TODO(), filter).Decode(&result)

	// Prints a message if no documents are matched or if any
	// other errors occur during the operation
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}

	// end findOne

	output, err := json.MarshalIndent(result, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n output:", output)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(output)
}
