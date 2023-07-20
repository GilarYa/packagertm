package packagertm

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataCS struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Namacs   string             `bson:"name,omitempty" json:"name,omitempty"`
	Emailcs  string             `bson:"email,omitempty" json:"email,omitempty"`
	Nohpcs   string             `bson:"nohp,omitempty" json:"nohp,omitempty"`
	Negaracs string             `bson:"negara,omitempty" json:"negara,omitempty"`
	Desccs   string             `bson:"desc,omitempty" json:"desc,omitempty"`
}
