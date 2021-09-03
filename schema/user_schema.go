package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Dados dos usuarios
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAT time.Time          `bson:"created_at"	json:"created_at"`
	UpdatedAT time.Time          `bson:"updated_at,omitempty"	json:"updated_at,omitempty"`
}

//Lista de usuarios
type Users []*User
