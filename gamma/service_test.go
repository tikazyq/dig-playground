package gamma

import (
	"dig-playground/alpha"
	"dig-playground/beta"
	"dig-playground/interfaces"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
	"testing"
)

func TestService_Run(t *testing.T) {
	var err error
	c := dig.New()
	err = c.Provide(alpha.NewService)
	require.Nil(t, err)
	err = c.Provide(beta.NewService)
	require.Nil(t, err)
	var alpha interfaces.AlphaService
	var beta interfaces.BetaService
	err = c.Invoke(func(alphaSvc interfaces.AlphaService, betaSvc interfaces.BetaService) {
		alpha = alphaSvc
		beta = betaSvc
	})
	require.Nil(t, err)
	svc, err := NewService(alpha, beta)
	require.Nil(t, err)
	svc.Run()
}
