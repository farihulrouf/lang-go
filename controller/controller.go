package controller

import (
	"fmt"
	"log"
    //"strings"
	"golangpro/helper"
	"golangpro/typedata"
	"go.mongodb.org/mongo-driver/bson"
)




func FindAllParnter() {
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
    fmt.Println(result)
    for csr.Next(helper.Ctx) {
        var row typedata.Partner
        err := csr.Decode(&row)
        if err != nil {
            log.Fatal(err.Error())
        }

        result = append(result, row)
        //fmt.Println(result)
    }
    fmt.Println("Codes:Name")
    fmt.Println("=======================")
    /*
    for _, c := range result {

        fmt.Println(c.Code,":", c.Name ,":", c.Enable)
    }
    */

}



func FindPartner() {
    db, err := helper.Connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    csr, err := db.Collection("rf_partners").Find(helper.Ctx, bson.M{"code": "GTX"})
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

    if len(result) > 0 {
        fmt.Println("Code", result[0].Code)
        fmt.Println("Name", result[0].Name)
        fmt.Println("Status: " ,result[0].Enable)
    }
}


