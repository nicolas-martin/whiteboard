package i

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type FakeDelegate struct {
	t         *testing.T
	expectedX int
	result    int
}

func (f *FakeDelegate) Do(x int) int {
	require.Equal(f.t, f.expectedX, x)
	return f.result
}

func TestAppFoo(t *testing.T) {
	app := &App{}
	app.delegate = &FakeDelegate{t, 1, 2}
	app.foo()
}
