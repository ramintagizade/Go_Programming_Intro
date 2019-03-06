package testers

import (
	"testing"

	"../calculate"
)

func TestCalculate(t *testing.T) {
	res, err := calculate.Calculate(4)

	if err != nil {
		t.Log("Error , wrong value  , expected value is ", res)
		t.Fail()
	}
}
