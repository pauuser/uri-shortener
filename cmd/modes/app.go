package modes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiber_recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"uri-shortener/cmd/modes/flags"
	"uri-shortener/internal/router/handlers"
	"uri-shortener/internal/services/repositories"
	"uri-shortener/internal/services/repositories/repositories_impl"
	"uri-shortener/internal/services/usecases"
	"uri-shortener/internal/services/usecases/usecases_impl"
)

type appRepoFields struct {
	linkRepository repositories.LinkRepository
}

type appUseCaseFields struct {
	linkUseCase usecases.LinkUseCase
}

type AppMode struct {
	app      *fiber.App
	Config   Config
	repos    appRepoFields
	useCases appUseCaseFields
}

type Config struct {
	Host  string           `mapstructure:"host"`
	Port  string           `mapstructure:"port"`
	Redis flags.RedisFlags `mapstructure:"redis"`
	Link  flags.LinkFlags  `mapstructure:"link_configuration"`
}

func (a *AppMode) ParseConfig(pathToConfig string, configFileName string) error {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetConfigType("json")
	v.AddConfigPath(pathToConfig)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&a.Config)
	if err != nil {
		return err
	}

	fmt.Println(a.Config)

	return nil
}

func (a *AppMode) Init() error {
	err := a.initRepositories()
	if err != nil {
		return err
	}
	err = a.initUsecases()
	if err != nil {
		return err
	}

	a.app = fiber.New()
	a.app.Use(fiber_recover.New())

	api := a.app.Group("")
	handlers.NewLinkHandler(api, a.useCases.linkUseCase)

	return nil
}

func (a *AppMode) initRepositories() error {
	repos := appRepoFields{}

	redisClient, err := a.Config.Redis.InitRedis()
	if err != nil {
		return err
	}

	repos.linkRepository = repositories_impl.NewLinkRepository(redisClient)
	a.repos = repos

	return nil
}

func (a *AppMode) initUsecases() error {
	services := appUseCaseFields{}

	services.linkUseCase = usecases_impl.NewLinkUseCase(a.repos.linkRepository, a.Config.Link)
	a.useCases = services

	return nil
}

func (a *AppMode) Run() {
	err := a.Init()
	if err == nil {
		err = a.app.Listen(a.Config.Host + ":" + a.Config.Port)
	}
}
