package service

import (
	"context"
	"fmt"

	"github.com/jack-hughes/users/internal/storage"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jack-hughes/users/internal/utils"
	"github.com/jack-hughes/users/pkg/api/users"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Storage defines the available functionality against the database. This interface is created so that the backend
// database technology can easily be swapped out, whilst maintaining the service contract
//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=./service.go -package=mocks -destination=../../test/mocks/storage_mocks.go
type Storage interface {
	Create(ctx context.Context, usr types.User) (types.User, error)
	Update(ctx context.Context, req types.User) (types.User, error)
	Delete(ctx context.Context, req types.User) error
	List(ctx context.Context, countryFilter string) ([]types.User, error)
}

// UserService contains the gRPC server definition alongside a logger,
// storage method interface, and a database connection pool
type UserService struct {
	users.UsersServer

	log   *zap.Logger
	store Storage
	db    storage.Postgres
}

// NewUserService instantiates the user service with a logger, database connection pool and storage interface
func NewUserService(log *zap.Logger, store Storage, db storage.Postgres) *UserService {
	return &UserService{
		log:   log.With(zap.String("component", "service")),
		store: store,
		db:    db,
	}
}

// Create allows the creation of new user objects in the database
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
	if err != nil {
		u.log.Error(fmt.Sprintf("failed to create user: %v", err))
		return &users.User{}, utils.SanitiseError(err)
	}

	u.log.Debug(fmt.Sprintf("created user with id: %v", usr.Id))
	return hydrateAPIResponse(usr), nil
}

// Update will update all present fields bar the user id and timestamps
// (which are managed by the database implementation)
func (u UserService) Update(ctx context.Context, user *users.User) (*users.User, error) {
	u.log.Debug(fmt.Sprintf("updating user with id: %v", user.Id))
	usr, err := u.store.Update(ctx, types.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	})
	if err != nil {
		u.log.Error(fmt.Sprintf("failed to update user: %v", err))
		return &users.User{}, utils.SanitiseError(err)
	}

	u.log.Debug(fmt.Sprintf("updated user with id: %v", usr.Id))
	return hydrateAPIResponse(usr), nil
}

// Delete will attempt to remove a record based on the user id presented
func (u UserService) Delete(ctx context.Context, user *users.User) (*empty.Empty, error) {
	u.log.Debug(fmt.Sprintf("deleting user with id: %v", user.Id))
	err := u.store.Delete(ctx, types.User{
		Id: user.Id,
	})
	if err != nil {
		u.log.Error(fmt.Sprintf("failed to delete user: %v", err))
		return &empty.Empty{}, utils.SanitiseError(err)
	}

	return &empty.Empty{}, nil
}

// List will return a slice of user types from the database, and return them on a stream back to the client
func (u UserService) List(req *users.ListUsersRequest, stream users.Users_ListServer) error {
	u.log.Debug(fmt.Sprintf("listing users for country code: %s", req.Filter))
	ctx := stream.Context()
	usrs, err := u.store.List(ctx, req.Filter)
	if err != nil {
		u.log.Error(fmt.Sprintf("failed to retrieve list of users: %v", err))
		return utils.SanitiseError(err)
	}

	u.log.Debug(fmt.Sprintf("number of users retrieved: %v", len(usrs)))
	for _, usr := range usrs {
		u.log.Debug(fmt.Sprintf("attempting to stream user: %v", usr.Id))
		if err := stream.Send(hydrateAPIResponse(usr)); err != nil {
			u.log.Error(fmt.Sprintf("failed to stream user: %v", usr.Id))
			return utils.SanitiseError(err)
		}
	}

	u.log.Debug("end of stream")

	return nil
}

// hydrateAPIResponse converts internal database representations of a user to the gRPC type to be returned
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
