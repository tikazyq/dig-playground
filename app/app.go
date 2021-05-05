package app

import (
	"dig-playground/alpha"
	"dig-playground/beta"
	"dig-playground/gamma"
	"dig-playground/interfaces"
	"go.uber.org/dig"
)

type App struct {
	alpha interfaces.AlphaService
	beta  interfaces.BetaService
	gamma interfaces.GammaService
}

func (app *App) Run() error {
	app.alpha.Init()
	app.beta.Setup()
	app.gamma.Run()
	app.alpha.Stop()
	app.beta.Cleanup()
	return nil
}

func NewApp() (*App, error) {
	app := &App{}
	c := dig.New()
	if err := c.Provide(alpha.NewService); err != nil {
		return nil, err
	}
	if err := c.Provide(beta.NewService); err != nil {
		return nil, err
	}
	if err := c.Provide(gamma.NewService); err != nil {
		return nil, err
	}
	if err := c.Invoke(func(alpha interfaces.AlphaService, beta interfaces.BetaService, gamma interfaces.GammaService) {
		app.alpha = alpha
		app.beta = beta
		app.gamma = gamma
	}); err != nil {
		return nil, err
	}
	return app, nil
}
