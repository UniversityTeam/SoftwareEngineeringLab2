package lab2

import (
	"testing"
	"strings"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	res, err := ConvertPostfixToPrefix("3 3 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "- 3 3", res)
	}
}

func TestPostfixToPrefixLabExample(t *testing.T) {
	res, err := ConvertPostfixToPrefix("4 2 - 3 * 5 +")
	assert.Equal(t, nil, err)
	assert.Equal(t, "+ * - 4 2 3 5", res)
}

func TestPostfixToPrefixOnlyPlus(t *testing.T) {
	res, err := ConvertPostfixToPrefix("4 17 + 3 + 51 + 19 + 21 + 78")
	assert.Equal(t, nil, err)
	assert.Equal(t, "+ + + + + 4 17 3 51 19 21 78", res)
}

func TestPostfixToPrefixOnlyMinus(t *testing.T) {
	res, err := ConvertPostfixToPrefix("4 3 - 13 - 21 - 5 - 19 - 1 - 2")
	assert.Equal(t, nil, err)
	assert.Equal(t, "- - - - - - 4 3 13 21 5 19 1 2", res)
}

func TestPostfixToPrefixOnlyMultiply(t *testing.T) {
	res, err := ConvertPostfixToPrefix("100 101 * 102 *")
	assert.Nil(t, err)
	assert.Equal(t, "* * 100 101 102", res)
}

func TestPostfixToPrefixOnlyDivide(t *testing.T) {
	res, err := ConvertPostfixToPrefix("18 2 / 3 /")
	assert.Nil(t, err)
	assert.Equal(t, "/ / 18 2 3", res)
}

func TestPostfixToPrefixAllOperators(t *testing.T) {
	res, err := ConvertPostfixToPrefix("54 6 3 * / 7 + 11 -")
	assert.Nil(t, err)
	assert.Equal(t, "- + / 54 * 6 3 7 11", res)
}

func TestPostfixToPrefixEmptyInput(t *testing.T) {
	res, err := ConvertPostfixToPrefix("")
	assert.Equal(t, "", res)
	assert.NotNil(t, err)
}

var res string

func BenchmarkPostfixToPrefix(b *testing.B) {
  expression := "54 6 3 * / 7 + 11 -"
  for i := 1; i <= 20; i++ {
    repeatedExpression := strings.Repeat(expression, i * i) + "10" //add 10 to avoid broking expression
    b.Run(fmt.Sprintf("%d-operators", i * i * 6), func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            res, _ = ConvertPostfixToPrefix(repeatedExpression)
        }
    })
  }
}