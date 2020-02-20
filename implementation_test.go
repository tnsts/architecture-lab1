package lab1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePostfix(t *testing.T) {
	res, err := PrefixToPostfix("1 2 + 3 * 4 - 5 /")
	if assert.Nil(t, err) {
		assert.Equal(t, 1, res)
	}
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 + 2 *")
	fmt.Println(res)

	// Output:
	// 8
}
