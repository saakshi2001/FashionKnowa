package main

import (
    "encoding/json"
    "fmt"
    "log"
    "context"
    "net/http"
	"html/template"
	"strings"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
type User struct {
    id  string `json:"id"`
    name string `json:"name"`
	email  string `json:"email"`
    password string `json:"password"`
}
func ShowUser(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    var id=r.Form["id"]
    var name=r.Form["name"]
    var email=r.Form["email"]
    var password=r.Form["password"]
    user := User{id,name,email,password}

    js, err := json.Marshal(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
	
}
func close(client *mongo.Client, ctx context.Context,cancel context.CancelFunc){
            
    defer cancel()
     
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()
}

func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error) {
 
    ctx, cancel := context.WithTimeout(context.Background(),
                                       30 * time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{})(*mongo.InsertOneResult, error) {
 
    collection := client.Database(dataBase).Collection(col)
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}

func main() {
 
    const port = 8000
  listenAt := fmt.Sprintf(":%d", port)
    http.HandleFunc("/users", ShowUser) 
  log.Printf("Open the following URL in the browser: http://localhost:%d\n", port)
  log.Fatal(http.ListenAndServe(listenAt, nil))

    client, ctx, cancel, err := connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer close(client, ctx, cancel)
    var user interface{}
    user = bson.D{
        {"id", id},
        {"name", name},
        {"email",email},
        {"password",password},
    }
    insertOneResult, err := insertOne(client, ctx, "insta","users", user)

    if err != nil {
        panic(err)
    }
}