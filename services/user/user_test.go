package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/phkress/mongo/model"
	repository "github.com/phkress/mongo/repositories/user"
)

func TestUserService_Create(t *testing.T) {
	validUser := model.User{
		ID:        "0xcv",
		Name:      "Glauber",
		Email:     "glauber.abigail@email.uk",
		CreatedAT: time.Time{},
		UpdatedAT: time.Time{},
	}

	validRepoMock := repository.NewMockUserRepository(gomock.NewController(t))
	validRepoMock.EXPECT().Create(validUser).Times(1).Return(nil)

	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		u       *UserService
		args    args
		wantErr bool
	}{
		{
			name: "Given a valid user model, it should create an user and return no errors",
			u: &UserService{
				repository: validRepoMock,
			},
			args: args{
				user: validUser,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
