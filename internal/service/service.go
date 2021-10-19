package service

import (
	"context"
	"github.com/v0ld-dev/etalon-arch/internal/domain"
	"github.com/v0ld-dev/etalon-arch/pkg/cache"
	"github.com/v0ld-dev/etalon-arch/pkg/otp"
	"github.com/v0ld-dev/etalon-arch/pkg/repository"
	"github.com/v0ld-dev/etalon-arch/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, userID primitive.ObjectID, hash string) error
	CreateSchool(ctx context.Context, userID primitive.ObjectID, schoolName string) (domain.School, error)
}

type Files interface {
	UploadAndSaveFile(ctx context.Context, file domain.File) (string, error)
	Save(ctx context.Context, file domain.File) (primitive.ObjectID, error)
	UpdateStatus(ctx context.Context, fileName string, status domain.FileStatus) error // TODO check schoolID
	GetByID(ctx context.Context, id, schoolId primitive.ObjectID) (domain.File, error)
	InitStorageUploaderWorkers(ctx context.Context)
}

type Services struct {
	Users	       Users
	Files          Files
}

type Deps struct {
	Repos                  *repository.Repositories
	Cache                  cache.Cache
	OtpGenerator           otp.Generator
	Environment            string
	StorageProvider        storage.Provider
}

func NewServices(deps Deps) *Services {

	usersService :=  NewUsersService(deps.Cache, deps.OtpGenerator)
	filesService :=  NewFilesService(deps.Repos.Files, deps.StorageProvider, deps.Environment, deps.Cache)

	return &Services{
		Users:    usersService,
		Files:    filesService,
	}
}

