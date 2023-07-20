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

func InsertDataCS(db *mongo.Database, namacs, emailcs, nohpcs string, negaracs string, desccs string) (InsertedID interface{}) {
	var dataCS DataCS
	dataCS.Namacs = namacs
	dataCS.Emailcs = emailcs
	dataCS.Nohpcs = nohpcs
	dataCS.Negaracs = negaracs
	dataCS.Desccs = desccs
	return InsertOneDoc(db, "data_DataCS", dataCS)
}

//	func GetDataNegara(negara string, db *mongo.Database, col string) (data DataCS) {
//		act := db.Collection(col)
//		filter := bson.M{"negara": negara}
//		err := act.FindOne(context.TODO(), filter).Decode(&data)
//		if err != nil {
//			fmt.Printf("getdataanegara: %v\n", err)
//		}
//		return data
//	}
func GetDataNamacs(namacs string, db *mongo.Database, col string) (data DataCS) {
	accou := db.Collection(col)
	filter := bson.M{"namacs": namacs}
	err := accou.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getNamacs: %v\n", err)
	}
	return data
}
func GetDataNegaracs(negaracs string, db *mongo.Database, col string) (data DataCS) {
	agd := db.Collection(col)
	filter := bson.M{"negaracs": negaracs}
	err := agd.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdatanegaracs: %v\n", err)
	}
	return data
}

// func DeleteDataNegara(negara string, db *mongo.Database, col string) (data DataCS) {
// 	dct := db.Collection(col)
// 	filter := bson.M{"negara": negara}
// 	err, _ := dct.DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf("DeleteDataCS : %v\n", err)
// 	}
// 	fmt.Println("Succes Delete data")
// 	return data
// }

func DeleteDataNamacs(namacs string, db *mongo.Database, col string) (data DataCS) {
	dena := db.Collection(col)
	filter := bson.M{"namacs": namacs}
	err, _ := dena.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataNama : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data

}
func DeleteDataNegaracs(negaracs string, db *mongo.Database, col string) (data DataCS) {
	dena := db.Collection(col)
	filter := bson.M{"negaracs": negaracs}
	err, _ := dena.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDatanegara : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
