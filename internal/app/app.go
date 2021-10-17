package app

import "github.com/v0ld-dev/etalon-arch/internal/service"

func Run() {
	memCache := cache.NewMemoryCache()

	services := service.NewServices(service.Deps{
		Cache: memCache,
	})
}
