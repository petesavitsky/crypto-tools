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

func (ethereum ethereumStruct) GetStringValue() string {
	return ethereum.stringValue
}

func (ethereum ethereumStruct) GetIntValue() int64 {
	return ethereum.intValue
}

func (ethereum ethereumStruct) Add(ethereumToAdd Ethereum) Ethereum {
	satoshis := ethereum.GetIntValue()
	satoshis += ethereumToAdd.GetIntValue()
	return NewEthereumFromInt(satoshis)
}

func (ethereum ethereumStruct) Subtract(ethereumToSubtract Ethereum) Ethereum {
	satoshis := ethereum.GetIntValue()
	satoshis -= ethereumToSubtract.GetIntValue()
	return NewEthereumFromInt(satoshis)
}

// Compare sort by amount ascending
func (ethereum ethereumStruct) Compare(other Ethereum) int {
	if ethereum.GetIntValue() > other.GetIntValue() {
		return 1
	} else if ethereum.GetIntValue() < other.GetIntValue() {
		return -1
	}
	return 0
}

func (ethereum ethereumStruct) GetCost(price USD) USD {
	cost := price.GetIntValue() * ethereum.GetIntValue() / ethCoinMultiplier
	return NewUSDFromInt(cost)
}

func (ethereum ethereumStruct) Multiply(value int64, fractionLength int64) Ethereum {
	percentMultiplier := int64(math.Pow10(int(fractionLength)))
	return NewEthereumFromInt(ethereum.GetIntValue() * value / percentMultiplier)
}

func (ethereum ethereumStruct) GetFractionLength() int64 {
	return ethIntFractionLength
}

//NewEthereumFromString create new ethereum based on string value
func NewEthereumFromString(ethString string) (Ethereum, error) {
	ethString = standardizeEthString(ethString)
	ethInt, err := ethStringToInt(ethString)
	if err != nil {
		return nil, err
	}
	ethereum := ethereumStruct{stringValue: ethString, intValue: ethInt}
	return ethereum, nil
}

//NewEthereumFromInt create a new ethereum based on int value
func NewEthereumFromInt(ethInt int64) Ethereum {
	ethString := ethIntToString(ethInt)
	return ethereumStruct{stringValue: ethString, intValue: ethInt}
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
		return 0, ConversionError{message: "Empty string passed for ethereum conversion [" + ethString + "]"}
	}
	pieces := strings.Split(ethString, ethSizeSeparator)
	fraction, err := convertEthFractionStringToInt(pieces[1])
	if err != nil {
		return 0, ConversionError{message: "Error converting fraction ethereum [" + pieces[1] + "] for string [" + ethString + "] -- [" + err.Error() + "]"}
	}
	wholeCoins, err := convertWholeEthToInt(pieces)
	if err != nil {
		return 0, ConversionError{message: "Error converting whole ethereum [" + pieces[0] + "] for string [" + ethString + "] -- [" + err.Error() + "]"}
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
