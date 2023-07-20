package packagertm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataCustomerServices(db *mongo.Database, nama, email, nohp string, negara string, desc string) (InsertedID interface{}) {
	var dataCustomerServices CustomerServices
	dataCustomerServices.Nama = nama
	dataCustomerServices.Email = email
	dataCustomerServices.Nohp = nohp
	dataCustomerServices.Negara = negara
	dataCustomerServices.Desc = desc
	return InsertOneDoc(db, "data_CustomerServices", dataCustomerServices)
}

func GetDataCustomerServices(negara string, db *mongo.Database, col string) (data CustomerServices) {
	act := db.Collection(col)
	filter := bson.M{"negara": negara}
	err := act.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataaccbyact: %v\n", err)
	}
	return data
}
func GetDataNama(nama string, db *mongo.Database, col string) (data CustomerServices) {
	accou := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := accou.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataCustomerService: %v\n", err)
	}
	return data
}
func DeleteDataCustomerServices(negara string, db *mongo.Database, col string) (data CustomerServices) {
	dct := db.Collection(col)
	filter := bson.M{"negara": negara}
	err, _ := dct.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataCustomerServices : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}

func DeleteDataNama(nama string, db *mongo.Database, col string) (data CustomerServices) {
	dena := db.Collection(col)
	filter := bson.M{"nama": nama}
	err, _ := dena.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataNama : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
