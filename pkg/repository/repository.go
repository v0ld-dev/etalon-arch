package repository

import (
	"context"
	"github.com/v0ld-dev/etalon-arch/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, email, password string) (domain.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.User, error)
	Verify(ctx context.Context, userID primitive.ObjectID, code string) error
	SetSession(ctx context.Context, userID primitive.ObjectID, session domain.Session) error
	AttachSchool(ctx context.Context, userID, schoolID primitive.ObjectID) error
}

type Files interface {
	Create(ctx context.Context, file domain.File) (primitive.ObjectID, error)
	UpdateStatus(ctx context.Context, fileName string, status domain.FileStatus) error
	GetForUploading(ctx context.Context) (domain.File, error)
	UpdateStatusAndSetURL(ctx context.Context, id primitive.ObjectID, url string) error
	GetByID(ctx context.Context, id, schoolID primitive.ObjectID) (domain.File, error)
}

type Repositories struct {
	Users          Users
	Files          Files
}
