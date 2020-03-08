package v1

import (
	context2 "golang.org/x/net/context"

	"github.com/nodias/golang.grpc.account.api/app/interface/rpc/v1.0/account"
	"github.com/nodias/golang.grpc.account.api/app/usecase"
)

type userService struct {
	userUsecase usecase.UserUsecase
}

func (s *userService) CreateUser(context2.Context, *account.CreateUserRequest) (*account.CreateUserResponse, error) {
	panic("implement me")
}

func (s *userService) ReadUsers(*account.ReadUsersRequest, account.AccountService_ReadUsersServer) error {
	panic("implement me")
}

func (s *userService) ReadUser(context2.Context, *account.ReadUserRequest) (*account.ReadUserResponse, error) {
	panic("implement me")
}

func (s *userService) UpdateUser(context2.Context, *account.UpdateUserRequest) (*account.UpdateUserResponse, error) {
	panic("implement me")
}

func (s *userService) DeleteUser(context2.Context, *account.DeleteUserRequest) (*account.DeleteUserResponse, error) {
	panic("implement me")
}

func NewUserService(userUsecase usecase.UserUsecase) *userService {
	return &userService{
		userUsecase: userUsecase,
	}
}


func toUser(users []*usecase.User) []*account.User {
	res := make([]*account.User, len(users))
	for i, user := range users {
		res[i] = &account.User{
			Id:    user.ID,
			Email: user.Email,
		}
	}
	return res
}
