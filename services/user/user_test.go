package user_service_test

import (
	"testing"
	"time"

	"github.com/phkress/mongo/models"
	userService "github.com/phkress/mongo/services/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := models.User{
		ID:        oid,
		Name:      "Glauber",
		Email:     "glauber.astrolino@email.com",
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	err := userService.Create(user)
	if err != nil {
		t.Error("teste de criacação de usuario falhou")
		t.Fail()
	} else {
		t.Log("teste de criação de usaurio com exito")
	}
}
func TestRead(t *testing.T) {

	users, err := userService.Read()

	if err != nil {
		t.Error("Erro - Consulta Usuario")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("Erro - Consulta não retornou Dados")
		t.Fail()
	} else {
		t.Log("Ok - Consulta Usuario")
	}

}
func TestUpdate(t *testing.T) {

	user := models.User{
		Email: "glauber.astrolino@email.com.uk",
	}

	err := userService.Update(user, userId)

	if err != nil {
		t.Error("Erro - Falha ao atualizar usuario")
		t.Fail()
	} else {
		t.Log("Ok - Usuario atualizado")
	}
}
func TestDelete(t *testing.T) {
	err := userService.Delete(userId)

	if err != nil {
		t.Error("Erro - Falaha ao excluir usuario")
		t.Fail()
	} else {
		t.Log("Ok - Usuario excluido com exito")
	}
}
