package assets

// Bitcoin bitcoin asset type
type Bitcoin interface {
	GetStringValue() string
	GetIntValue() int64
	Add(Bitcoin) Bitcoin
	Subtract(Bitcoin) Bitcoin
	GetCost(USD) USD
	Multiply(value int64, fractionDigits int64) Bitcoin
	GetFractionLength() int64
	Compare(Bitcoin) int
	GetUnitCostAtPrice(USD) USD
}

type bitcoinStruct struct {
	stringValue string
	intValue    int64
}

// USD asset type
type USD interface {
	GetStringValue() string
	GetPrettyStringValue() string
	GetIntValue() int64
	Add(USD) USD
	Subtract(USD) USD
	Multiply(value int64, fractionDigits int64) USD
	Compare(USD) int
	GetFractionLength() int64
}

type usdStruct struct {
	stringValue string
	intValue    int64
	usdStringFractionLength      int64
	usdIntFractionLength         int64
	usdFractionSignificantDigits int64
}

// Ether  asset type
type Ether interface {
	GetStringValue() string
	GetIntValue() int64
	Add(Ether) Ether
	Subtract(Ether) Ether
	GetCost(USD) USD
	Multiply(value int64, percentMultiplier int64) Ether
	GetFractionLength() int64
	Compare(Ether) int
}

type etherStruct struct {
	stringValue string
	intValue    int64
}

// ConversionError error converting
type ConversionError struct {
	message string
}

func (err ConversionError) Error() string {
	return err.message
}
