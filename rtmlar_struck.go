package packagertm

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataCS struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama   string             `bson:"name,omitempty" json:"name,omitempty"`
	Email  string             `bson:"email,omitempty" json:"email,omitempty"`
	Nohp   string             `bson:"nohp,omitempty" json:"nohp,omitempty"`
	Negara string             `bson:"negara,omitempty" json:"negara,omitempty"`
	Desc   string             `bson:"desc,omitempty" json:"desc,omitempty"`
}
