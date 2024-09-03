package adapters

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dario-labs/srv/configuration"
	"github.com/dario-labs/srv/domain/pokemon"
	"github.com/dario-labs/srv/domain/repositories"
	v1 "github.com/dario-labs/srv/handler/v1"
	pokerepo "github.com/dario-labs/srv/repository/pokemon"
	pokemonSorter "github.com/dario-labs/srv/repository/pokemon/sorter"
	"github.com/dario-labs/srv/repository/shared"
	"github.com/dario-labs/srv/usecases"
	"github.com/gin-gonic/gin"
)

type Params struct {
	Page     int
	PerPage  int
	Criteria string
}

func (s *Server) routes() {
	s.router = gin.New()
	core := s.router.Group("/")
	core.Use()

	// compose service logic
	conf := configuration.GetConfiguration()

	pokeAPIClient := shared.CreateClient(conf.Host, conf.APIKey)

	// TODO: info resource is the same (a.k.a. poke api), this should simulate Gen1Repo as a different resource from others gens.
	pokemonGen1Repo := pokerepo.NewPokemonsRepository(pokeAPIClient)

	pokemonRepositoryMap := map[pokemon.PokemonGenIDPrefixType]repositories.PokemonRepository{
		pokemon.Gen1: pokemonGen1Repo,
	}

	// PokeAPI doesn't have sorting engine, handling a custom one.
	pkmSorter := pokemonSorter.NewPokemonSorter()

	pokemonUseCases := usecases.NewPokemonUseCases(pokemonRepositoryMap, pkmSorter)

	pokemonService := v1.PokemonService{
		UseCases: pokemonUseCases,
	}

	// Routes
	{
		// R1: listing Pokemón info, applying sorting and filtering criteria
		core.GET("/pokemons", func(c *gin.Context) { //should move logic to a handler before calling the service
			var params Params

			pageStr := c.Query("page")
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				log.Fatalln("invalid page value")
				return
			}
			perPageStr := c.Query("per_page")
			perPage, err := strconv.Atoi(perPageStr)
			if err != nil {
				log.Fatalln("invalid per page value")
				return
			}

			params = Params{
				Page:     page,
				PerPage:  perPage,
				Criteria: c.Query("criteria"),
			}

			data, err := pokemonService.GetPokemons(c, v1.Params(params))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
				return
			}

			c.IndentedJSON(http.StatusOK, data)
		})

		// R2: getting a pokemon info
		core.GET("/pokemon/:poke_id", func(c *gin.Context) { //should move logic to a handler before calling the service
			var err error
			pkmID, err := strconv.ParseInt(c.Param("poke_id"), 10, 32)
			if err != nil || pkmID < 1 {
				log.Fatalln("Pokemon ID should be a positive integer")
				return
			}

			data, err := pokemonService.GetPokemonDetail(c, int(pkmID))
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, err)
			}

			c.IndentedJSON(http.StatusOK, data)
		})
	}
}
