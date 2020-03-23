package assets

import "testing"

func TestBitcoin(t *testing.T) {
  btc, _ := NewBitcoinFromString("1.20500282")
  usd, _ := NewUSDFromString("7055.38")
  expectedUnitCost, _ := NewUSDFromString("5855.07")
  actualUnitCost := btc.GetUnitCostAtPrice(usd)
  if expectedUnitCost.Compare(actualUnitCost) == 0 {
    t.Log("Great success!")
  } else {
    t.Errorf("Failure. Expected %s but got %s", expectedUnitCost.GetStringValue(), actualUnitCost.GetStringValue())
  }
}

func TestNegativeBitcoin(t *testing.T) {
  btc, err := NewBitcoinFromString("-1.00")
  if err != nil {
    t.Error(err)
  }
  if btc.GetStringValue() != "-1.00000000" {
    t.Errorf("Invalid btc string value returned %s", btc.GetStringValue())
  }

}
