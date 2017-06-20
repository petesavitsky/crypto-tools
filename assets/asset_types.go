package assets

// Bitcoin bitcoin asset type
type Bitcoin interface {
	GetStringValue() string
	GetIntValue() int64
	Add(Bitcoin) Bitcoin
	Subtract(Bitcoin) Bitcoin
	GetCost(USD) USD
	Multiply(value int64, percentMultiplier int64) Bitcoin
	GetFractionLength() int64
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
	Multiply(value int64, percentMultiplier int64) USD
	Compare(USD) int
	GetFractionLength() int64
}

type usdStruct struct {
	stringValue string
	intValue    int64
}

// Ethereum  asset type
type Ethereum interface {
	GetStringValue() string
	GetIntValue() int64
	Add(Ethereum) Ethereum
	Subtract(Ethereum) Ethereum
	GetCost(USD) USD
	Multiply(value int64, percentMultiplier int64) Ethereum
	GetFractionLength() int64
}

type ethereumStruct struct {
	stringValue string
	intValue    int64
}

type ConversionError struct {
	message string
}

func (err ConversionError) Error() string {
	return err.message
}
