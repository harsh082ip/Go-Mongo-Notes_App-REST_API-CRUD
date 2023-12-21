package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func GetNotes(w http.ResponseWriter, r *http.Request) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
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

	coll := client.Database("sample_restaurants").Collection("restaurants")
	filter := bson.D{{}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	fmt.Println("Cursor: ", cursor)
	for cursor.Next(context.TODO()) {
		var result Notes
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}

		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", output)
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)
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

	coll := client.Database("sample_restaurants").Collection("restaurants")

	objectID, err := primitive.ObjectIDFromHex(id["id"])
	if err != nil {
		panic(err)
	}

	filter := bson.D{{"_id", objectID}}
	var deletedDocument Notes
	err = coll.FindOneAndDelete(context.TODO(), filter).Decode(&deletedDocument)
	if err != nil {
		panic(err)
	}

	// Print the deleted document
	output, err := bson.MarshalExtJSON(deletedDocument, false, false)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deleted Document:\n%s\n", output)

	// var delNote Notes
	// _ = json.NewDecoder(deleteResult).Decode(delNote)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	var note Notes
	err = json.Unmarshal(body, &note)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	note.Id = primitive.NewObjectID()

	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Error Disconnect from MongoDB: ", err)
		}
	}()

	// make connection with MongoDB
	coll := client.Database("sample_restaurants").Collection("restaurants")

	// convert Go struct to BSON document
	document, err := bson.Marshal(note)
	if err != nil {
		log.Fatal(err)
	}
	_, err = coll.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Inserted Doc: ", document)
	fmt.Println("Inserted Doc: ", note)
	fmt.Fprintf(w, "Inserted Document Successfully :")
	jsonNote, err := json.Marshal(note)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonNote)
}
