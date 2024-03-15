package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimal(t *testing.T) {
	num1 := 29.444444444
	num2 := 29.555555555
	assert.Equal(t, float64(29), Decimal(num1))
	assert.Equal(t, float64(30), Decimal(num2))
}
