package stringcalculator_test

import (
	"testing"

	sc "github.com/jeremybaumont/stringcalculator-kata/stringcalculator"
	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	assert.Equal(t, 42, sc.Add("42"))
}
