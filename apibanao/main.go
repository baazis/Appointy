package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/baazis/appointy/helper"
	"github.com/baazis/appointy/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var collection = helper.ConnectDB()

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route

	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")

	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")

	// r.HandleFunc("/posts/users/{id}", getUserPosts).Methods("GET")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getUser(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user models.Users
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)

	

}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.Users

	// we decode our body request params
	// hashedPassword, err := HashPassword(user.Password)
	hashedPassword, err := HashPassword(user.Password)

	
	_ = json.NewDecoder(r.Body).Decode(&user)

	// insert our user model.
	user.Password = string(hashedPassword)
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)

	// Next, insert the username, along with the hashed password into the database
	// if _, err = db.Query("insert into users values ($1, $2)", user.Name, string(hashedPassword)); err != nil {
	// 	// If there is any issue with inserting into the database, return a 500 error
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	json.NewEncoder(w).Encode(result)
	// json.NewEncoder(w).Encode(hashedPassword)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var post models.Post
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Post

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&post)

	// insert our user model.
	result, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// func getUserPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	//var post models.Posts
// 	var user models.Users
// 	var params = mux.Vars(r)

// 	id, _ := primitive.ObjectIDFromHex(params["id"])

// 	filter := bson.M{"_id": id}
// 	err := collection.FindOne(context.TODO(), filter).Decode(&user)
// 	if err != nil {
// 		helper.GetError(err, w)
// 		return
// 	}

// 	fmt.Println("This route is under construction :)")
// 	json.NewEncoder(w).Encode(user.Posts)
// }

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
