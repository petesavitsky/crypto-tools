package assets

import (
	"bytes"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	btcSizeSeparator        = "."
	btcStringFractionLength = 8
	btcIntFractionLength    = 8
)

var btcCoinMultiplier = int64(math.Pow10(btcIntFractionLength))

func (bitcoin bitcoinStruct) GetStringValue() string {
	return bitcoin.stringValue
}

func (bitcoin bitcoinStruct) GetIntValue() int64 {
	return bitcoin.intValue
}

func (bitcoin bitcoinStruct) Add(bitcoinToAdd Bitcoin) Bitcoin {
	satoshis := bitcoin.GetIntValue()
	satoshis += bitcoinToAdd.GetIntValue()
	return NewBitcoinFromInt(satoshis)
}

func (bitcoin bitcoinStruct) Subtract(bitcoinToSubtract Bitcoin) Bitcoin {
	satoshis := bitcoin.GetIntValue()
	satoshis -= bitcoinToSubtract.GetIntValue()
	return NewBitcoinFromInt(satoshis)
}

func (bitcoin bitcoinStruct) GetCost(price USD) USD {
	cost := price.GetIntValue() * bitcoin.GetIntValue() / btcCoinMultiplier
	return NewUSDFromInt(cost)
}

func (bitcoin bitcoinStruct) Multiply(value int64, fractionLength int64) Bitcoin {
	percentMultiplier := int64(math.Pow10(int(fractionLength)))
	return NewBitcoinFromInt(bitcoin.GetIntValue() * value / percentMultiplier)
}

func (bitcoin bitcoinStruct) GetFractionLength() int64 {
	return btcIntFractionLength
}

// Compare sort by amount ascending
func (bitcoin bitcoinStruct) Compare(other Bitcoin) int {
	if bitcoin.GetIntValue() > other.GetIntValue() {
		return 1
	} else if bitcoin.GetIntValue() < other.GetIntValue() {
		return -1
	}
	return 0
}

func (bitcoin bitcoinStruct) GetUnitCostAtPrice(price USD) USD {
	ratio := float64(price.GetIntValue()) / float64(bitcoin.intValue)
	multiplier := math.Pow10(int(btcIntFractionLength))
	unitCostFloat := ratio * multiplier
	unitCostInt := int64(math.Round(unitCostFloat))
	return NewUSDFromInt(unitCostInt)
}

//NewBitcoinFromString create new bitcoin based on string value
func NewBitcoinFromString(btcString string) (Bitcoin, error) {
	btcString = standardizeBtcString(btcString)
	btcInt, err := btcStringToInt(btcString)
	if err != nil {
		return nil, err
	}
	bitcoin := bitcoinStruct{stringValue: btcString, intValue: btcInt}
	return bitcoin, nil
}

//NewBitcoinFromInt create a new bitcoin based on int value
func NewBitcoinFromInt(btcInt int64) Bitcoin {
	btcString := btcIntToString(btcInt)
	return bitcoinStruct{stringValue: btcString, intValue: btcInt}
}

func standardizeBtcString(btcString string) string {
	pieces := strings.Split(btcString, btcSizeSeparator)
	fractionString := "0"
	if len(pieces) > 1 {
		fractionString = pieces[1]
		fractionLength := len(fractionString)
		if fractionLength > ethIntFractionLength {
			fractionString = fractionString[0:btcIntFractionLength]
		}
	}
	fractionLength := utf8.RuneCountInString(fractionString)
	var fractionBuffer bytes.Buffer
	fractionBuffer.WriteString(fractionString)
	for i := fractionLength; i < btcStringFractionLength; i++ {
		fractionBuffer.WriteString("0")
	}
	var standardizedBuffer bytes.Buffer
	standardizedBuffer.WriteString(pieces[0])
	standardizedBuffer.WriteString(btcSizeSeparator)
	standardizedBuffer.WriteString(fractionBuffer.String())
	return standardizedBuffer.String()
}

func btcStringToInt(btcString string) (int64, error) {
	if len(btcString) < 1 {
		return 0, ConversionError{message: "Empty string passed for bitcoin conversion [" + btcString + "]"}
	}
	pieces := strings.Split(btcString, btcSizeSeparator)
	fraction, err := convertBtcFractionStringToInt(pieces[1])
	if err != nil {
		return 0, ConversionError{message: "Error converting fraction bitcoin [" + pieces[1] + "] for string [" + btcString + "] -- [" + err.Error() + "]"}
	}
	wholeCoins, err := convertWholeBtcToInt(pieces)
	if err != nil {
		return 0, ConversionError{message: "Error converting whole bitcoin [" + pieces[0] + "] for string [" + btcString + "] -- [" + err.Error() + "]"}
	}
	return wholeCoins + fraction, nil
}

// add padding to get to btcFractionLength decimals then convert to int64
func convertBtcFractionStringToInt(fractionString string) (int64, error) {
	fraction, err := strconv.Atoi(fractionString)
	if err != nil {
		return 0, err
	}
	return int64(fraction), nil
}

func convertWholeBtcToInt(pieces []string) (int64, error) {
	if pieces[0] == "" {
		return 0, nil
	}
	wholeCoins, err := strconv.Atoi(pieces[0])
	if err != nil {
		return 0, err
	}
	return int64(wholeCoins) * btcCoinMultiplier, nil
}

func btcIntToString(btcInt int64) string {
	wholeBtc := btcInt / int64(btcCoinMultiplier)
	fractionBtc := btcInt % int64(btcCoinMultiplier)
	fractionString := strconv.FormatInt(fractionBtc, 10)
	fractionLength := utf8.RuneCountInString(fractionString)
	var buffer bytes.Buffer
	for i := fractionLength; i < btcStringFractionLength; i++ {
		buffer.WriteString("0")
	}
	btcString := strconv.FormatInt(wholeBtc, 10) + btcSizeSeparator + buffer.String() + fractionString
	return btcString
}
