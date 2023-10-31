package iparam

import (
	"errors"
	"testing"
)

type simpleMock struct {
	err error
}

func (m *simpleMock) DoSomething() error {
	return m.err
}

func TestDoSomething(t *testing.T) {
	errorMock := &simpleMock{errors.New("some error")}
	Hello(errorMock)
	// test that "an error occurred" is logged

	regularMock := &simpleMock{}
	Hello(regularMock)
	// test "no error occurred" is logged
}
