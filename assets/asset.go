package assets

import "math"

// DivideByAsset divides an asset by another, outputing in the specified fraction decimal length
func DivideByAsset(dividend Asset, divisor Asset, outputFractionDecimalLength int64) Asset {
  ratio := float64(dividend.GetIntValue()) / float64(divisor.GetIntValue())
  multiplierPower := (divisor.GetFractionLength() - dividend.GetFractionLength()) + outputFractionDecimalLength
  multiplier := math.Pow10(int(multiplierPower))
  resultFloat := math.Round(ratio * multiplier)
  resultInt := int64(resultFloat)
  return assetStruct{value: resultInt, fractionLength: outputFractionDecimalLength}
}

// Divide divids an assent by an int, outputs an asset with the specified fraction decimal length
func Divide(dividend Asset, divisor, outputFractionDecimalLength int64) Asset {
  ratio := float64(dividend.GetIntValue()) / float64(divisor)
  multiplierPower := outputFractionDecimalLength - dividend.GetFractionLength()
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
