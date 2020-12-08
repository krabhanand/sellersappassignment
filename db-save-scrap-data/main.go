package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	var p ProductData
	w.Header().Set("Content-Type", "application/json")

	//get mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://anand:anand@cluster0.rlk7s.mongodb.net/"))

	// handle client error
	if err != nil {
		errorsend := Error{
			ErrorInfo: err.Error(),
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorsend)
		return
	}

	//connect to client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//handle connection error
	if err != nil {
		errorsend := Error{
			ErrorInfo: err.Error(),
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorsend)
		return
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	merr := json.NewDecoder(r.Body).Decode(&p)

	//input decoding error, bad request
	if merr != nil {
		http.Error(w, merr.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Connected to MongoDB!")

	//get collection from required database
	collection := client.Database("seller-app").Collection("product-data")

	//update timestamp
	p.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	//search if document exists
	var podcast ProductData
	if err = collection.FindOne(ctx, bson.M{"url": p.URL}).Decode(&podcast); err != nil {

		//if document is not found inset document
		insertResult, err := collection.InsertOne(context.TODO(), p)

		//handle errors of insertion
		if err != nil {
			errorsend := Error{
				ErrorInfo: err.Error(),
			}
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(errorsend)
			return
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	} else {

		//update the document
		resultupdate, uerr := collection.ReplaceOne(
			ctx,
			bson.M{"url": p.URL},
			p,
		)

		//handle updation error
		if uerr != nil {
			errorsend := Error{
				ErrorInfo: err.Error(),
			}
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(errorsend)
			return
		}

		fmt.Printf("Replaced %v Documents!\n", resultupdate)

	}

	// close the connection
	err = client.Disconnect(context.TODO())

	//handle connection closing error
	if err != nil {
		errorsend := Error{
			ErrorInfo: err.Error(),
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorsend)
		return
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}

	//send json response
	json.NewEncoder(w).Encode(p)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Product struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Price       string `json:"price" bson:"price,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Reviews     string `json:"reviews" bson:"reviews,omitempty"`
	ImageURL    string `json:"imageurl" bson:"imageurl,omitempty"`
}

type ProductData struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	URL       string              `json:"url" bson:"url,omitempty"`
	Product   *Product            `json:"product" bson:"product,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
type Error struct {
	ErrorInfo string `json:"error_info,omitempty" bson:"error_info,omitempty"`
}
