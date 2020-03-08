package repository

import (
	"github.com/nodias/golang.grpc.account.api/app/domain/model"
)


//TODO Maybe this package should be delete, because of ORM of mongodb
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(*model.User) error
}
