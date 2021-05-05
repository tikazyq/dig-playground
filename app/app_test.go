package app

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApp_Run(t *testing.T) {
	app, err := NewApp()
	require.Nil(t, err)
	err = app.Run()
	require.Nil(t, err)
}
