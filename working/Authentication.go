package work

import (
	"context"
	data "demo/Data"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Collection, context.Context, *mongo.Collection, *mongo.Client) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//y
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Access the "test" database and "users" collection
	database := client.Database("database")
	collection := database.Collection("Clients")
	admincollection := database.Collection("Admins")
	return collection, ctx, admincollection, client
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}
	collection, ctx, _, client := Connect()
	w.Header().Set("Content-Type", "application/json")
	var db data.User
	var user data.User
	json.NewDecoder(r.Body).Decode(&db)
	filter := bson.D{{Key: "userid", Value: db.Userid}}
	defer client.Disconnect(ctx)
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {

			http.Error(w, "Database error", http.StatusInternalServerError)
		}

		return
	}
	if r.URL.Path != "/login" {
		http.Error(w, "Path error", http.StatusNotFound)
		return
	}
	userpass := user.Password
	dbpass := db.Password
	if userpass != dbpass {
		w.Write([]byte("Incorrect password"))
		return
	}
	if db.Userid == user.Userid && userpass == dbpass {
		fmt.Fprint(w, "Welcome ", user.Name)

	}
	Gettip(w, r, user.Technology)

}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Please free to Register")
	collection, ctx, _, client := Connect()
	var user data.User
	var check data.User
	json.NewDecoder(r.Body).Decode(&user)

	filter := bson.D{{Key: "userid", Value: user.Userid}}
	err := collection.FindOne(ctx, filter).Decode(&check)
	defer client.Disconnect(ctx)
	if err == nil {
		http.Error(w, "Userid not Available", http.StatusNotAcceptable)
		return

	}
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "data not stored in database", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(&result)
	w.Write([]byte("Registered  successfully"))
}

func Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}
	_, ctx, admincollection, client := Connect()
	w.Header().Set("Content-Type", "application/json")
	var user data.User
	var db data.User
	json.NewDecoder(r.Body).Decode(&db)
	filter := bson.D{{Key: "userid", Value: db.Userid}}
	defer client.Disconnect(ctx)
	err := admincollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {

			http.Error(w, "Database error", http.StatusInternalServerError)
		}

		return
	}
	if r.URL.Path != "/admin" {
		http.Error(w, "Path error", http.StatusNotFound)
		return
	}
	userpass := user.Password
	dbpass := db.Password
	if userpass != dbpass {
		w.Write([]byte("Incorrect password"))
		return
	}
	if db.Userid == user.Userid && userpass == dbpass {
		w.Write([]byte("Sigin succesfull"))
		AddTip(w, r)
	}

}
func Gettip(w http.ResponseWriter, r *http.Request, tech string) {
	now := time.Now()
	//w.Header().Set("Content-Type", "application/json")
	// Extract the day of the month (date)
	day := now.Day()
	var user1 data.Tip
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Tool := client.Database("Tools")
	techh := Tool.Collection(tech)
	filter := bson.D{{Key: "date", Value: day}}
	err = techh.FindOne(ctx, filter).Decode(&user1)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
	}
	jsonBytes, err := json.Marshal(user1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

}
func AddTip(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not Supported", http.StatusMethodNotAllowed)
		return
	}

	_, ctx, _, client := Connect()
	w.Header().Set("Content-Type", "application/json")
	defer client.Disconnect(ctx)
	// Parse JSON data from the request body into a Tip struct
	var tip data.Tip
	if err := json.NewDecoder(r.Body).Decode(&tip); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database := client.Database("Tools")
	godata := database.Collection(tip.Technology)
	_, err := godata.InsertOne(ctx, tip)

	if err != nil {
		http.Error(w, "data not found", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Tip added successfully"))
}
