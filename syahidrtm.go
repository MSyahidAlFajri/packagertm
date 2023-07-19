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
func InsertDataRTM(db *mongo.Database, nama, tanggal string, lokasi string, agenda string) (InsertedID interface{}) {
	var datartm DataRTM
	datartm.NamaRapat = nama
	datartm.TanggalRapat = tanggal
	datartm.LokasiRapat = lokasi
	datartm.AgendaRapat = agenda
	return InsertOneDoc(db, "data_rtm", datartm)
}

func GetDataRtmFromAgenda(agendarapat string, db *mongo.Database, col string) (data DataRTM) {
	agd := db.Collection(col)
	filter := bson.M{"agendarapat": agendarapat}
	err := agd.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdatartmbyagd: %v\n", err)
	}
	return data
}

func GetDataRtm(namarapat string, db *mongo.Database, col string) (data []DataRTM) {
	rapat := db.Collection(col)
	filter := bson.M{"namarapat": namarapat}
	cursor, err := rapat.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataRtm: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteDataRtm(lokasirapat string, db *mongo.Database, col string) (data DataRTM) {
	lks := db.Collection(col)
	filter := bson.M{"lokasirapat": lokasirapat}
	err, _ := lks.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataRtm : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
