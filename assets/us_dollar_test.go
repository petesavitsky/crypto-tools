package assets

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	dollarsUno, err := NewUSDFromString("0.00307472")
	if err != nil {
		t.Errorf("Error parsing dollars uno %v", err)
		return
	}
	t.Logf("Uno - %s -- int %d -- frac %d", dollarsUno.GetStringValue(), dollarsUno.GetIntValue(), dollarsUno.GetFractionLength())
	dollarsOne, err := NewUSDFromString("18755.77")
	if err != nil {
		t.Errorf("Error parsing dollars one %v", err)
		return
	}
	expectedDollars, err := NewUSDFromString("450136.42")
	if err != nil {
		t.Errorf("Error parsing expected dollars %v", err)
		return
	}
	dollarsThree := dollarsOne.Multiply(2399989, 5)
	if !(dollarsThree.Compare(expectedDollars) == 0) {
		t.Errorf("Dollars are not equal -- actual dollars %s", dollarsThree.GetStringValue())
	} else {
		t.Log("Dollars are equal!")
	}
}

func TestSubtraction(t *testing.T) {
	higher, _ := NewUSDFromString("100.00")
	lower, _ := NewUSDFromString("50.00")
	result := lower.Subtract(higher)
	if "-50.00" != result.GetStringValue() {
		t.Errorf("Invalid subtraction result %s", result.GetStringValue())
	}
}

func TestNegative(t *testing.T) {
	dollars, err := NewUSDFromString("-34.50")
	if err != nil {
		t.Errorf("couldn't parse negative %s", dollars.GetStringValue())
	}
	negDollars := NewUSDFromInt(-1000)
	if "-10.00" != negDollars.GetStringValue() {
		t.Errorf("Invalid negative string value %s", negDollars.GetStringValue())
	}
}
