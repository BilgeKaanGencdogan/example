package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NewContractKeyPrefix is the prefix to retrieve all NewContract
	NewContractKeyPrefix = "NewContract/value/"
)

// NewContractKey returns the store key to retrieve a NewContract from the index fields
// basically this function creates a byte type key of "NewContract/value/{dealId}/"
func NewContractKey(
	index string,
) []byte {
	var key []byte
	key = append(key, []byte(NewContractKeyPrefix)...)
	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)
	return key
}
