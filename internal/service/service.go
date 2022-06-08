package service

import (
	"context"
	"fmt"
	"github.com/jack-hughes/users/internal/storage"
	"github.com/jack-hughes/users/internal/utils"
	"github.com/jack-hughes/users/pkg/api/users"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	users.UsersServer

	log   *zap.Logger
	store storage.Storage
}

func NewUserService(log *zap.Logger, db storage.Storage) *UserService {
	return &UserService{
		log:   log.With(zap.String("component", "service")),
		store: db,
	}
}

func (u UserService) Create(ctx context.Context, user *users.User) (*users.User, error) {
	u.log.Debug(fmt.Sprintf("creating user with email: %v", user.Email))

	usr, err := u.store.Create(ctx, user)
	return hydrateAPIResponse(usr), utils.SanitiseError(err)
}

func (u UserService) Update(ctx context.Context, user *users.User) (*users.User, error) {
	return &users.User{}, nil
}

func (u UserService) Delete(ctx context.Context, user *users.User) (*users.User, error) {
	return &users.User{}, nil
}

func (u UserService) List(req *users.ListUsersRequest, stream users.Users_ListServer) error {
	ctx := stream.Context()
	usrs, err := u.store.List(ctx, req)
	if err != nil {
		return utils.SanitiseError(err)
	}

	for _, usr := range usrs {
		if err := stream.Send(hydrateAPIResponse(usr)); err != nil {
			return utils.SanitiseError(err)
		}
	}

	return nil
}

func hydrateAPIResponse(in storage.User) *users.User {
	return &users.User{
		Id:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Nickname:  in.Nickname,
		Password:  in.Password,
		Email:     in.Email,
		Country:   in.Country,
		CreatedAt: timestamppb.New(in.CreatedAt),
		UpdatedAt: timestamppb.New(in.UpdatedAt),
	}
}
