package packagertm

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "DataCS",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	Namacs := "Gilar"
	Emailcs := "gilarya@gmail.com"
	Nohpcs := "085312345678"
	Negaracs := "Indonesia"
	Desccs := "Ada yang bisa di bantu?"
	hasil := InsertDataCS(MongoConn, Namacs, Emailcs, Nohpcs, Negaracs, Desccs)
	fmt.Println(hasil)

}

func TestGetDataDataCS(t *testing.T) {
	Namacs := "Gilar"
	hasil := GetDataCS(Namacs, MongoConn, "data_DataCS")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Negaracs := "Indonesia"
	hasil := DeleteDataCS(Negaracs, MongoConn, "data_DataCS")
	fmt.Println(hasil)

}
