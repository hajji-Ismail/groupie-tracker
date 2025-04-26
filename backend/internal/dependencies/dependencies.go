package dependencies

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/services"
)

type Dependencies struct {
	// Handlers
	AllArtistsHandler    *handlers.AllArtistsHandler
	ArtistDetailsHandler *handlers.ArtistsDetailsHandler
	FilterHandler        *handlers.FilterHandler
	SearchHandler               *handlers.SearchHandler
	// UserHandler *handlers.UserHandler
}

func NewDependencies() *Dependencies {
	// load data in first
	store := repository.New_Store()
	store.LoadData()

	// Instantiate services
	allartistservice := services.NewAllArtistsService(store)
	artistDetailsservice := services.NewArtistDetailsService(store)
	filterservice := services.NewFilterService(store)
	searchService := services.NewSearchService(store)

	// Instantiate handlers
	return &Dependencies{
		// ItemService: itemService,
		// UserService:  userService,
		AllArtistsHandler:    handlers.NewAllArtistsHandler(allartistservice),
		ArtistDetailsHandler: handlers.NewArtistDetailsService(artistDetailsservice),
		FilterHandler:        handlers.NewFilterHandler(filterservice),
		SearchHandler:               handlers.NewSearchHandler(searchService),
		// UserHandler:  userHandler,
	}
}
