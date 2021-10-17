package service

import "github.com/v0ld-dev/etalon-arch/pkg/cache"


type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, userID primitive.ObjectID, hash string) error
	CreateSchool(ctx context.Context, userID primitive.ObjectID, schoolName string) (domain.School, error)
}

type Services struct {
	Users	Users
}

type Deps struct {
	Cache                  cache.Cache
}

func NewServices(deps Deps) *Services {

	usersService :=  NewUsersService(deps.Cache)

	return &Services{
		Users:    usersService,
	}
}

