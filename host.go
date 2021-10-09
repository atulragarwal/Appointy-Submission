package main

//Atul Agarwal
//19BCE0436

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "regexp"
)

type User struct {
	ID       int    `bson:"id" json:"id"`
	Name     string `bson:"name" json:name`
	Email    string `bson:"email" json:email`
	Password string `bson:"password" json:password`
}

type Post struct {
	ID        int       `bson:"id,omitempty"`
	PostId    int       `bson:"postId,omitempty"`
	Caption   string    `bson:"caption,omitempty"`
	ImageUrl  string    `bson:"imageUrl,omitempty"`
	TimeStamp time.Time `bson:"timeStamp,omitempty"`
}

func CheckUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost:
		MakeUser(w, r)
		return
	case r.Method == http.MethodGet:
		CheckUser(w, r)
		return
	}
}

func MakeUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tempPassword := user.Password
	passwordEncoded := base64.StdEncoding.EncodeToString([]byte(tempPassword))
	user.Password = passwordEncoded
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("user")
	enterResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", enterResult.InsertedID)
	json.NewEncoder(w).Encode(enterResult.InsertedID)
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("user")
	var episodes []User
	tempId := r.URL.Query().Get("id")
	newTemp, err := strconv.Atoi(tempId)
	cursor, err := collection.Find(ctx, bson.M{"id": newTemp})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	var tempUser User = episodes[0]
	storePass := "Encoded Password: " + tempUser.Password
	tempPass, err := base64.StdEncoding.DecodeString(tempUser.Password)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	tempUser.Password = string(tempPass)
	episodes[0] = tempUser
	fmt.Println(episodes)
	json.NewEncoder(w).Encode(storePass)
	json.NewEncoder(w).Encode(episodes)
}
func CheckPostUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost:
		MakePost(w, r)
		return
	case r.Method == http.MethodGet:
		CheckPost(w, r)
		return
	}
}

func MakePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.TimeStamp = time.Now()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("post")
	enterResult, err := collection.InsertOne(ctx, post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", enterResult.InsertedID)
	json.NewEncoder(w).Encode(enterResult.InsertedID)
}

func CheckPost(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("post")
	var episodes []Post
	tempId := r.URL.Query().Get("id")
	newTemp, err := strconv.Atoi(tempId)
	cursor, err := collection.Find(ctx, bson.M{"postId": newTemp})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes)
	json.NewEncoder(w).Encode(episodes)
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("post")
	var episodes []Post
	tempId := r.URL.Query().Get("id")
	newTemp, err := strconv.Atoi(tempId)
	cursor, err := collection.Find(ctx, bson.M{"id": newTemp})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes)
	json.NewEncoder(w).Encode(episodes)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", CheckUrl)
	mux.HandleFunc("/posts/", CheckPostUrl)
	mux.HandleFunc("/posts/users/", GetUserPosts)
	http.ListenAndServe(":8080", mux)

}
