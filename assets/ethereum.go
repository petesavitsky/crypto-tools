package assets

import (
	"bytes"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	ethSizeSeparator        = "."
	ethStringFractionLength = 6
	ethIntFractionLength    = 6
)

var ethCoinMultiplier = int64(math.Pow10(ethIntFractionLength))

func (ether etherStruct) GetStringValue() string {
	return ether.stringValue
}

func (ether etherStruct) GetIntValue() int64 {
	return ether.intValue
}

func (ether etherStruct) Add(etherToAdd Ether) Ether {
	satoshis := ether.GetIntValue()
	satoshis += etherToAdd.GetIntValue()
	return NewEtherFromInt(satoshis)
}

func (ether etherStruct) Subtract(etherToSubtract Ether) Ether {
	satoshis := ether.GetIntValue()
	satoshis -= etherToSubtract.GetIntValue()
	return NewEtherFromInt(satoshis)
}

// Compare sort by amount ascending
func (ether etherStruct) Compare(other Ether) int {
	if ether.GetIntValue() > other.GetIntValue() {
		return 1
	} else if ether.GetIntValue() < other.GetIntValue() {
		return -1
	}
	return 0
}

func (ether etherStruct) GetCost(price USD) USD {
	cost := price.GetIntValue() * ether.GetIntValue() / ethCoinMultiplier
	return NewUSDFromInt(cost)
}

func (ether etherStruct) Multiply(value int64, fractionLength int64) Ether {
	percentMultiplier := int64(math.Pow10(int(fractionLength)))
	return NewEtherFromInt(ether.GetIntValue() * value / percentMultiplier)
}

func (ether etherStruct) GetFractionLength() int64 {
	return ethIntFractionLength
}

//NewEtherFromString create new ether based on string value
func NewEtherFromString(ethString string) (Ether, error) {
	ethString = standardizeEthString(ethString)
	ethInt, err := ethStringToInt(ethString)
	if err != nil {
		return nil, err
	}
	ether := etherStruct{stringValue: ethString, intValue: ethInt}
	return ether, nil
}

//NewEtherFromInt create a new ether based on int value
func NewEtherFromInt(ethInt int64) Ether {
	ethString := ethIntToString(ethInt)
	return etherStruct{stringValue: ethString, intValue: ethInt}
}

func standardizeEthString(ethString string) string {
	pieces := strings.Split(ethString, ethSizeSeparator)
	fractionString := "0"
	if len(pieces) > 1 {
		fractionString = pieces[1]
	}
	fractionLength := utf8.RuneCountInString(fractionString)
	var fractionBuffer bytes.Buffer
	fractionBuffer.WriteString(fractionString)
	for i := fractionLength; i < ethStringFractionLength; i++ {
		fractionBuffer.WriteString("0")
	}
	var standardizedBuffer bytes.Buffer
	standardizedBuffer.WriteString(pieces[0])
	standardizedBuffer.WriteString(ethSizeSeparator)
	standardizedBuffer.WriteString(fractionBuffer.String())
	return standardizedBuffer.String()
}

func ethStringToInt(ethString string) (int64, error) {
	if len(ethString) < 1 {
		return 0, ConversionError{message: "Empty string passed for ether conversion [" + ethString + "]"}
	}
	pieces := strings.Split(ethString, ethSizeSeparator)
	fraction, err := convertEthFractionStringToInt(pieces[1])
	if err != nil {
		return 0, ConversionError{message: "Error converting fraction ether [" + pieces[1] + "] for string [" + ethString + "] -- [" + err.Error() + "]"}
	}
	wholeCoins, err := convertWholeEthToInt(pieces)
	if err != nil {
		return 0, ConversionError{message: "Error converting whole ether [" + pieces[0] + "] for string [" + ethString + "] -- [" + err.Error() + "]"}
	}
	return wholeCoins + fraction, nil
}

// add padding to get to ethFractionLength decimals then convert to int64
func convertEthFractionStringToInt(fractionString string) (int64, error) {
	fraction, err := strconv.Atoi(fractionString)
	if err != nil {
		return 0, err
	}
	return int64(fraction), nil
}

func convertWholeEthToInt(pieces []string) (int64, error) {
	if pieces[0] == "" {
		return 0, nil
	}
	wholeCoins, err := strconv.Atoi(pieces[0])
	if err != nil {
		return 0, err
	}
	return int64(wholeCoins) * ethCoinMultiplier, nil
}

func ethIntToString(ethInt int64) string {
	wholeEth := ethInt / int64(ethCoinMultiplier)
	fractionEth := ethInt % int64(ethCoinMultiplier)
	fractionString := strconv.FormatInt(fractionEth, 10)
	fractionLength := utf8.RuneCountInString(fractionString)
	var buffer bytes.Buffer
	for i := fractionLength; i < ethStringFractionLength; i++ {
		buffer.WriteString("0")
	}
	ethString := strconv.FormatInt(wholeEth, 10) + ethSizeSeparator + buffer.String() + fractionString
	return ethString
}
