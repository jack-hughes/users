package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jack-hughes/users/internal/storage"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jack-hughes/users/internal/utils"
	"github.com/jack-hughes/users/pkg/api/users"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=./service.go -package=mocks -destination=../../test/mocks/storage_mocks.go
type Storage interface {
	Create(ctx context.Context, usr types.User) (types.User, error)
	Update(ctx context.Context, req types.User) (types.User, error)
	Delete(ctx context.Context, req types.User) error
	List(ctx context.Context, countryFilter string) ([]types.User, error)
}

type UserService struct {
	users.UsersServer

	log   *zap.Logger
	store Storage
	db    storage.Postgres
}

func NewUserService(log *zap.Logger, store Storage, db storage.Postgres) *UserService {
	return &UserService{
		log:   log.With(zap.String("component", "service")),
		store: store,
		db:    db,
	}
}

func (u UserService) Create(ctx context.Context, user *users.User) (*users.User, error) {
	u.log.Debug(fmt.Sprintf("creating user with email: %v", user.Email))

	usr, err := u.store.Create(ctx, types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	})
	return hydrateAPIResponse(usr), utils.SanitiseError(err)
}

func (u UserService) Update(ctx context.Context, user *users.User) (*users.User, error) {
	usr, err := u.store.Update(ctx, types.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	})
	return hydrateAPIResponse(usr), utils.SanitiseError(err)
}

func (u UserService) Delete(ctx context.Context, user *users.User) (*empty.Empty, error) {
	err := u.store.Delete(ctx, types.User{
		Id: user.Id,
	})

	return &empty.Empty{}, utils.SanitiseError(err)
}

func (u UserService) List(req *users.ListUsersRequest, stream users.Users_ListServer) error {
	ctx := stream.Context()
	usrs, err := u.store.List(ctx, req.Filter)
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

func hydrateAPIResponse(in types.User) *users.User {
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
