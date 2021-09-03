package model

import (
	"time"
)

//Dados dos usuarios
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAT time.Time
	UpdatedAT time.Time
}

//Lista de usuarios
type Users []*User
