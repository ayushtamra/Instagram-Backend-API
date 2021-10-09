package main

import (
	"encoding/json"
	"context"
	"fmt"
	"log"
	"time"
	"strings"
	// "math/rand"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "github.com/joho/godotenv"
)

var clientOptions = options.Client().ApplyURI("mongodb+srv://ayush:PASSWORD@cluster0.cgolj.mongodb.net/insta_backend_api?retryWrites=true&w=majority")
var client, err = mongo.Connect(context.TODO(), clientOptions)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

type Post struct {
	Userid   string `json:"uid"`
	Id       string `json:"id"`
	Caption  string `json:"caption"`
	ImageURL string `json:"imgurl"`
	Time     string `json:"time"`
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "mainpage endpoint hit, get request acknowledged")
}

func add_new_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u) // parsing json to structure
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := client.Database("insta_backend_api").Collection("users")
	insertResult, err1 := collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err1)
	}
	json.NewEncoder(w).Encode(insertResult) // response in json
}

func add_new_post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var p Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := client.Database("insta_backend_api").Collection("posts")
	insertResult, err1 := collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatal(err1)
	}
	json.NewEncoder(w).Encode(insertResult)
}

func find_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var found_user User
	q := r.URL.String()
	split := strings.Split(q, "/")
	prequery := split[len(split)-1] //to get the user ID from the URL
	query := bson.D{{"id", prequery}}
	collection := client.Database("insta_backend_api").Collection("users")
	err = collection.FindOne(context.TODO(), query).Decode(&found_user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(found_user) //response in json
}

func find_post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var found_post Post
	q := r.URL.String()
	split := strings.Split(q, "/")
	prequery := split[len(split)-1] //to get the post ID from the URL
	query := bson.D{{"id", prequery}}
	collection := client.Database("insta_backend_api").Collection("posts")
	err = collection.FindOne(context.TODO(), query).Decode(&found_post)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(found_post) // response in json
}

func find_user_post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var found_user_post []*Post
	q := r.URL.String()
	split := strings.Split(q, "/")
	prequery := split[len(split)-1] //to get the user ID from the URL
	query := bson.D{{"userid", prequery}}
	//limit, _ := strconv.ParseInt(split[len(split)-2],10,64)
	//offset, _ := strconv.ParseInt(split[len(split)-1],10,64)
	collection := client.Database("insta_backend_api").Collection("posts")
	//findOptions := options.Find()
	//findOptions.SetLimit(limit)
	//findOptions.SetSkip(offset)
	cursor, err2 := collection.Find(context.TODO(), query)
	if err2 != nil {
		log.Fatal(err2)
	}
	for cursor.Next(context.TODO()) {
		var temp_post Post // iterating cursor and adding decoding each found document into struct from json
		err := cursor.Decode(&temp_post)
		if err != nil {
			log.Fatal(err)
		}
		found_user_post = append(found_user_post, &temp_post)
	}
	json.NewEncoder(w).Encode(found_user_post) // response in json
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Password secure using Environment Variables
	// err := godotenv.Load(".env")
	// if err != nil {
	//   log.Fatal("Error loading .env file")
	// }
	// admin_pass := os.Getenv("ADMIN_PASSWORD")


	//MongoDB connect
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ayush:qujNnfBY7UQj9T2C@cluster0.cgolj.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)


	// Route Handlers / Endpoints
	// r.HandleFunc("/api/books", getBooks).Methods("GET");
	// r.HandleFunc("/api/books/{id}", getBook).Methods("GET");
	// r.HandleFunc("/api/books", createBook).Methods("POST");
	// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT");
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE");
	
	http.HandleFunc("/", mainpage)
	r.HandleFunc("/users", add_new_user).Methods("POST")
	r.HandleFunc("/users/{id}", find_user).Methods("GET")
	r.HandleFunc("/posts", add_new_post).Methods("POST")
	r.HandleFunc("/posts/{id}", find_post).Methods("GET")
	r.HandleFunc("/posts/users/{id}", find_user_post).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}