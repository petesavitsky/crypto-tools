package assets

import (
  "testing"
)

func TestMultiply(t *testing.T) {
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
