package modes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiber_recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

type AppMode struct {
	app    *fiber.App
	Config Config
}

type Config struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
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
	a.app = fiber.New()
	a.app.Use(fiber_recover.New())

	// apiV1 := a.app.Group("/api/v1")

	return nil
}

func (a *AppMode) Run() {
	err := a.Init()
	if err != nil {
		a.app.Listen(a.Config.Host + ":" + a.Config.Port)
	}
}
