package app

import (
	"github.com/v0ld-dev/etalon-arch/internal/service"
	"github.com/v0ld-dev/etalon-arch/pkg/cache"
)

func Run() {
	memCache := cache.NewMemoryCache()

	services := service.NewServices(service.Deps{
		Cache: memCache,
	})
}
