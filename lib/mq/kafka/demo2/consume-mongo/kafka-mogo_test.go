package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
	"testing"
)

func TestK_changeData(t*testing.T)  {
	var mySlice = []byte{0, 0, 0, 0, 0, 0, 0, 23}
	data := binary.BigEndian.Uint64(mySlice)
	fmt.Println(data)
}

func TestK_getData(t*testing.T)  {
	mongoURL := "mongodb://admin:1234qwer@localhost:27017" //os.Getenv("mongoURL")
	dbName := "example_db"//os.Getenv("dbName")
	collectionName := "example_coll2"//os.Getenv("collectionName")
	collection := getMongoCollection(mongoURL, dbName, collectionName)

	v:= collection.FindOne(context.Background(), bson.D{{"offset", 137}}, )
	log.Printf("info:: %+v\n",v)

}
