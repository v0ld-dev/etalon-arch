package service

import (
	"context"
	"github.com/v0ld-dev/etalon-arch/pkg/cache"
	"github.com/v0ld-dev/etalon-arch/pkg/otp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService struct {
	otpGenerator otp.Generator
	cache cache.Cache
}

func NewUsersService(cache cache.Cache, otpGenerator otp.Generator) *UsersService {
	return &UsersService{
		otpGenerator: otpGenerator,
		cache:        cache,
	}
}


func (u *UsersService) SignUp(ctx context.Context, input UserSignUpInput) error {
	panic("implement me")
}

func (u *UsersService) SignIn(ctx context.Context, input UserSignInInput) (interface{}, error) {
	panic("implement me")
}

func (u *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (interface{}, error) {
	panic("implement me")
}

func (u *UsersService) Verify(ctx context.Context, userID primitive.ObjectID, hash string) error {
	panic("implement me")
}

func (u *UsersService) CreateSchool(ctx context.Context, userID primitive.ObjectID, schoolName string) (interface{}, error) {
	panic("implement me")
}
