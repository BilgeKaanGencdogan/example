package types

const (
	// ModuleName defines the module name
	ModuleName = "deal"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_deal"
)

var (
	ParamsKey = []byte("p_deal")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DealCounterKey = "DealCounter-value-"
)

const (
	ContractCounterKey = "ContractCounter-value-"
)

// In types/keys.go
func NewDealStoreKey(dealId string) []byte {
	return append(KeyPrefix(NewDealKeyPrefix), NewDealKey(dealId)...)
}
