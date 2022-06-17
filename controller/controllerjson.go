package controller

import (
	"log"
    "strings"
    "net/http"
	"golangpro/helper"
	"encoding/json"
	"golangpro/typedata"
	"go.mongodb.org/mongo-driver/bson"
)

/*

    err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
        [
            {
                "$project": {
                    "exchange": 1,
                    "sell": {
                        "$cond": [
                            {
                                $eq: [
                                    "$transactiontype",
                                    "S"
                                ]
                            },
                            1,
                            0
                        ]
                    }
                }
            },
            {
                "$group": {
                    "_id" "$exchange",
                    "sell": {
                        "$sum": "$sell"
                    }
                }
            },
            {
              "$project": {
                "_id": false,
                "sell": 1,
                "exchange": "$_id"
              }  
            },
            {
               "$sort": {
                "sell": -1
               } 
            }
        ]
    `)), true, &pipeline)
    if err != nil {
        log.Fatal(err.Error())
    }

    csr, err := db.Collection("UploadTTReport").Aggregate(helper.Ctx, pipeline)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer csr.Close(helper.Ctx)

    result := make([]typedata.Selling, 0)
    for csr.Next(helper.Ctx) {
        var row typedata.Selling
        err := csr.Decode(&row)
        if err != nil {
            log.Fatal(err.Error())
        }

        result = append(result, row)
    }
    if r.Method == "GET" {
        var dataresult, err = json.Marshal(result)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Write(dataresult)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
    
}*/


func AggregateJson(w http.ResponseWriter, r *http.Request) {
    db, err := helper.Connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    pipeline := make([]bson.M, 0)

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
    if r.Method == "GET" {
        var dataresult, err = json.Marshal(result)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Write(dataresult)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
    
}

func FindAllParnterJson(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")
    db, err := helper.Connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    csr, err := db.Collection("rf_partners").Find(helper.Ctx, bson.M{})
    if err != nil {
        log.Fatal(err.Error())
    }
    defer csr.Close(helper.Ctx)

    result := make([]typedata.Partner, 0)
    for csr.Next(helper.Ctx) {
        var row typedata.Partner
        err := csr.Decode(&row)
        if err != nil {
            log.Fatal(err.Error())
        }

        result = append(result, row)
    }


    if r.Method == "GET" {
    	var dataresult, err = json.Marshal(result)

    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
            return
    	}
    	w.Write(dataresult)
    	return
    }

    http.Error(w, "", http.StatusBadRequest)

}

func FindUploadReportJson(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Content-Type", "application/json")
    db, err := helper.Connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    csr, err := db.Collection("UploadTTReport").Find(helper.Ctx, bson.M{})
    if err != nil {
        log.Fatal(err.Error())
    }
    defer csr.Close(helper.Ctx)

    result := make([]typedata.UploadTTReport, 0)
    for csr.Next(helper.Ctx) {
        var row typedata.UploadTTReport
        err := csr.Decode(&row)
        if err != nil {
            log.Fatal(err.Error())
        }

        result = append(result, row)
    }


    if r.Method == "GET" {
        var dataresult, err = json.Marshal(result)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Write(dataresult)
        return
    }

    http.Error(w, "", http.StatusBadRequest)

}




