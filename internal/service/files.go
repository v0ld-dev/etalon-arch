package service

import (
	"context"
	"github.com/v0ld-dev/etalon-arch/internal/domain"
	"github.com/v0ld-dev/etalon-arch/pkg/cache"
	"github.com/v0ld-dev/etalon-arch/pkg/repository"
	"github.com/v0ld-dev/etalon-arch/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilesService struct {
	repo    repository.Files
	storage storage.Provider
	env     string
	cache   cache.Cache
}


func NewFilesService(repo repository.Files, storage storage.Provider, env string, cache cache.Cache) *FilesService {
	return &FilesService{repo: repo, storage: storage, env: env, cache: cache}
}


func (f *FilesService) UploadAndSaveFile(ctx context.Context, file domain.File) (string, error) {
	panic("implement me")
}

func (f *FilesService) Save(ctx context.Context, file domain.File) (primitive.ObjectID, error) {
	panic("implement me")
}

func (f *FilesService) UpdateStatus(ctx context.Context, fileName string, status domain.FileStatus) error {
	panic("implement me")
}

func (f *FilesService) GetByID(ctx context.Context, id, schoolId primitive.ObjectID) (domain.File, error) {
	panic("implement me")
}

func (f *FilesService) InitStorageUploaderWorkers(ctx context.Context) {
	panic("implement me")
}


