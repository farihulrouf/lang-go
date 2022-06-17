
package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = func() context.Context {
	return context.Background()
}()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}



func connect() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("l_golang"), nil
}
/*
func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Silver", 4})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Gold", 5})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}
*/

func main () {
	//insert()
	//find()
	//update()
	//remove()
	aggregate()
}

func aggregate() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{ "$group": {
				"_id": null,
				"Total": { "$sum": 1 }
			} },
			{ "$project": {
				"Total": 1,
				"_id": 0
			} }
		]
	`)), true, &pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Total :", result[0]["Total"])
	}
}


/*func remove() {
    db, err := connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    var selector = bson.M{"name": "ethan"}
    _, err = db.Collection("student").DeleteOne(ctx, selector)
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Println("Remove success!")
}
*/

/*
func update() {
    db, err := connect()
    if err != nil {z
        log.Fatal(err.Error())
    }

    var selector = bson.M{"name": "John Wick"}
    var changes = student{"Farhul Rouf", 2}
    _, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Println("Update success!")
}
*/

/*
func find() {
    db, err := connect()
    if err != nil {
        log.Fatal(err.Error())
    }

	//db.getCollection("student").find({"name": "Wick"})

// mongo-go-driver
	//db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
    csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
    if err != nil {
        log.Fatal(err.Error())
    }
    defer csr.Close(ctx)

    result := make([]student, 0)
    for csr.Next(ctx) {
        var row student
        err := csr.Decode(&row)
        if err != nil {
            log.Fatal(err.Error())
        }

        result = append(result, row)
    }

    if len(result) > 0 {
        fmt.Println("Name  :", result[0].Name)
        fmt.Println("Grade :", result[0].Grade)
    }
}
*/
