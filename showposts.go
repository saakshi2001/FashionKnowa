package main

import (
    "context" 
    "fmt"     
    "os"
    "reflect" 
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Posts struct {
    id  string
    caption string
    url  string
	timestamp string
}

func main(){
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        fmt.Println("mongo.Connect() ERROR:", err)
        os.Exit(1)
    }
    ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
    col := client.Database("users").Collection("posts")
    fmt.Println("Posts:", reflect.TypeOf(col), "\n")
    var result Posts

    err = col.FindOne(context.TODO(), bson.D{}).Decode(&result)
    if err != nil {
        fmt.Println("FindOne() ERROR:", err)
        os.Exit(1)
    } else {
        fmt.Println("id:", result.id)
        fmt.Println("Caption:", result.caption)
        fmt.Println("URL:", result.url)
		fmt.Println("Posted Timestamp:", result.timestamp)
    }
    cursor, err := col.Find(context.TODO(), bson.D{})
    if err != nil {
        fmt.Println("Finding all documents ERROR:", err)
        defer cursor.Close(ctx)
    } else {

        for cursor.Next(ctx) {

            var result bson.M
            err := cursor.Decode(&result)
            if err != nil {
                fmt.Println("cursor.Next() error:", err)
                os.Exit(1)
            } else {
                fmt.Println("\nresult type:", reflect.TypeOf(result))
                fmt.Println("result:", result)
            }
        }
    }
}