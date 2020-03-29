package assets

import "math"

func Divide(dividend Asset, divisor Asset, outputFractionDecimalLength int64) Asset {
  ratio := float64(dividend.GetIntValue()) / float64(divisor.GetIntValue())
  multiplierPower := (divisor.GetFractionLength() - dividend.GetFractionLength()) + outputFractionDecimalLength
  multiplier := math.Pow10(int(multiplierPower))
  resultFloat := math.Round(ratio * multiplier)
  resultInt := int64(resultFloat)
  return assetStruct{value: resultInt, fractionLength: outputFractionDecimalLength}
}


func (asset assetStruct) GetIntValue() int64 {
  return asset.value
}

func (asset assetStruct) GetFractionLength() int64 {
  return asset.fractionLength
}
