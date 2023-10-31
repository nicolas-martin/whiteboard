package f

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppFoo(t *testing.T) {
	app := &App{}
	app.delegate = func(x int) int {
		require.Equal(t, 1, x)
		return 2
	}
	app.foo()
}
