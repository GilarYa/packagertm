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
	DBName:   "CustomerService",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	Nama := "Gilar"
	Email := "gilarya@gmail.com"
	Nohp := "085312345678"
	Negara := "Indonesia"
	Desc := "Ada yang bisa di bantu?"
	hasil := InsertDataCustomerServices(MongoConn, Nama, Email, Nohp, Negara, Desc)
	fmt.Println(hasil)

}

func TestGetDataCustomerServices(t *testing.T) {
	Nama := "Gilar"
	hasil := GetDataCustomerServices(Nama, MongoConn, "data_CustomerServices")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Negara := "Indonesia"
	hasil := DeleteDataCustomerServices(Negara, MongoConn, "data_CustomerServices")
	fmt.Println(hasil)

}
