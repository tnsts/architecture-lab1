package lab1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePostfix(t *testing.T) {

	tests := []struct {
		input string
		exp int
	}{
		{"1 2 + 3 * 4 - 5 /", 1},
		{"2 3 ^ 1 +", 9},
		{"7 3 + 5 / 2 4 ^ * 8 - 6 + 1 +", 3},
		{"4 2 - 3 * 5 +", 11},
		{"6 2 ^ 9 / 1 * 2 3 * - 3 + 8 +", 9},
	}

	for _, test := range tests{
		res, err := CalculatePostfix(test.input)
		if assert.Nil(t, err) {
			assert.Equal(t, res, test.exp)
		}
	}

	_, err := CalculatePostfix("15&*")
	assert.NotNil(t, err)

	_, err1 := CalculatePostfix("")
	assert.NotNil(t, err1)

}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 + 2 *")
	fmt.Println(res)

	// Output:
	// 8
}
