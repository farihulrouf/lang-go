package controller

import (
	"fmt"
	"log"
    "strings"
	"golangpro/helper"
	"golangpro/typedata"
	"go.mongodb.org/mongo-driver/bson"
    "encoding/json"
)


func Aggregate() {
    db, err := helper.Connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    pipeline := make([]bson.M, 0)
    /*
      { "$group": 
             {
                "_id": "$code",
                "Total": { "$sum": 1 }
             } },
             { "$project": {
                "Total": 1,
                
             } }
    */
    err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
        [
            { "$group": {
                "_id": {
                    "code" : "$code"
                },
                "Total": { "$sum": 1 }
            } }
        ]
    `)), true, &pipeline)
    if err != nil {
        log.Fatal(err.Error())
    }

    csr, err := db.Collection("rf_partners").Aggregate(helper.Ctx, pipeline)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer csr.Close(helper.Ctx)

    result := make([]typedata.AgaPrtner, 0)
    for csr.Next(helper.Ctx) {
        var row typedata.AgaPrtner
        err := csr.Decode(&row)
        if err != nil {
            log.Fatal(err.Error())
        }

        result = append(result, row)
    }
    fmt.Println("Codes:Name")
    fmt.Println("=======================")
    for _, c := range result {

        fmt.Println("Code:" ,c.Id.Code)
        fmt.Println("Total:", c.Total)
    }

    var jsonData, errot = json.Marshal(result)
    if err != nil {
        log.Fatal(errot.Error())
    }
    var jsonString = string(jsonData)
    fmt.Println(jsonString)

    /*result := make([]bson.M, 0)
    for csr.Next(helper.Ctx) {
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
    */
}







