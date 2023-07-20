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
	Nama := "Gilar"
	Email := "gilarya@gmail.com"
	Nohp := "085312345678"
	Negara := "Indonesia"
	Desc := "Ada yang bisa di bantu?"
	hasil := InsertDataDataCS(MongoConn, Nama, Email, Nohp, Negara, Desc)
	fmt.Println(hasil)

}

func TestGetDataDataCS(t *testing.T) {
	Nama := "Gilar"
	hasil := GetDataDataCS(Nama, MongoConn, "data_DataCS")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Negara := "Indonesia"
	hasil := DeleteDataDataCS(Negara, MongoConn, "data_DataCS")
	fmt.Println(hasil)

}
