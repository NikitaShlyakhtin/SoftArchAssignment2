package main

import (
	"github.com/stretchr/testify/require"
	"testing"

	"go.uber.org/fx"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(getFxOptions())
	require.NoError(t, err)
}
