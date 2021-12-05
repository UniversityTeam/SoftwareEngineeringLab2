package lab2

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

type OutputSim struct {
	called bool
}

func (o *OutputSim) Write(p []byte) (n int, err error) {
	o.called = true
	return 0, nil
}

func TestComputerHandleCorrect(t *testing.T) {
	testExpression := "4 2 - 3 * 5 +"
	inputSimulation := strings.NewReader(testExpression)
	outputSimulation := OutputSim{}
	handler := ComputeHandler{
		Input:  inputSimulation,
		Output: &outputSimulation,
	}
	err := handler.Compute()
	assert.Nil(t, err)
	assert.Equal(t, true, outputSimulation.called)
}

func TestComputerHandleIncorrect(t *testing.T) {
	testExpression := ""
	inputSimulation := strings.NewReader(testExpression)
	outputSimulation := OutputSim{}
	handler := ComputeHandler{
		Input:  inputSimulation,
		Output: &outputSimulation,
	}
	err := handler.Compute()
	assert.NotNil(t, err)
}