package main

import (
    "log"
    "net/http"
    "os"
	"context"
    "fmt"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context,
           cancel context.CancelFunc) {
    defer cancel()
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()
}
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
    ctx, cancel := context.WithTimeout(context.Background(),30 * time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}
func query(client *mongo.Client, ctx context.Context,dataBase, col string, query, field interface{})(result *mongo.Cursor, err error) {
    collection := client.Database(dataBase).Collection(col)
    result, err = collection.Find(ctx, query,options.Find().SetProjection(field))
    return
}

func main() {

client, ctx, cancel, err := connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer close(client, ctx, cancel)
 r, err := http.Get("/posts/:id")

    if err != nil {
      log.Fatal(err)
    }

    defer r.Body.Close()  
	 var filter, option interface{}
    filter = bson.D{
        {"id", bson.D{{"$eq", id}}},
    }

    if err != nil {
      log.Fatal(err)
    }

    defer r.Body.Close()

    f, err := os.Create("getpost.html")

    if err != nil {
      log.Fatal(err)
    }

    defer f.Close()

    _, err = f.ReadFrom(r.Body)

    if err != nil {
      log.Fatal(err)
    }
}