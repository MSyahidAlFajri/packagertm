package packagertm

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataRTM struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaRapat    string             `bson:"namarapat,omitempty" json:"namarapat,omitempty"`
	TanggalRapat string             `bson:"tanggalrapat,omitempty" json:"tanggalrapat,omitempty"`
	LokasiRapat  string             `bson:"lokasirapat,omitempty" json:"lokasirapat,omitempty"`
	AgendaRapat  string             `bson:"agendarapat,omitempty" json:"agendarapat,omitempty"`
}
