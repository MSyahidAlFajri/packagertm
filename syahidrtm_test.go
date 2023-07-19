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
	DBName:   "RTM",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	Nama := "Rapat Akreditasi"
	Tanggal := "25 Agustus 2023"
	Lokasi := "Auditorium"
	Agenda := "Membahas Untuk Akreditasi Jurusan Teknik Informatika"
	hasil := InsertDataRTM(MongoConn, Nama, Tanggal, Lokasi, Agenda)
	fmt.Println(hasil)

}

func TestGetDataRtm(t *testing.T) {
	NamaRapat := "Rapat Akreditasi"
	hasil := GetDataRtm(NamaRapat, MongoConn, "data_rtm")
	fmt.Println(hasil)

}

func TestGetDataRtmbyAgenda(t *testing.T) {
	Agenda := "Membahas Untuk Akreditasi Jurusan Teknik Informatika"
	hasil := GetDataRtmFromAgenda(Agenda, MongoConn, "data_rtm")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	lokasi := "Auditorium"
	hasil := DeleteDataRtm(lokasi, MongoConn, "data_rtm")
	fmt.Println(hasil)

}
