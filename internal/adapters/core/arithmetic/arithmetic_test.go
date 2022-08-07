package arithmetic_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.hexagonal-architecture/internal/adapters/core/arithmetic"
)

func TestAdditional(t *testing.T) {
	arith := arithmetic.NewAdapter()

	actual, _ := arith.Additional(3, 2)

	require.Equal(t, int32(5), actual)
}

func TestSubtraction(t *testing.T) {
	arith := arithmetic.NewAdapter()

	actual, _ := arith.Subtraction(3, 2)

	require.Equal(t, int32(1), actual)
}

func TestMultiplication(t *testing.T) {
	arith := arithmetic.NewAdapter()

	actual, _ := arith.Multiplication(3, 2)

	require.Equal(t, int32(6), actual)
}

func TestDivision(t *testing.T) {
	arith := arithmetic.NewAdapter()

	actual, _ := arith.Division(3, 2)

	require.Equal(t, int32(1), actual)
}
