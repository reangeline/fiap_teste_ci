package internal_test

import (
	"testing"

	"github.com/reangeline/fiap_teste_ci/internal"
	"github.com/stretchr/testify/assert"
)

func TestSomar(t *testing.T) {

	resultado, err := internal.Somar(10, 10)

	assert.NoError(t, err, resultado)

}
