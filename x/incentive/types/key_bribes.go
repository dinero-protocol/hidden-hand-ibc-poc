package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BribesKeyPrefix is the prefix to retrieve all Bribes
	BribesKeyPrefix = "Bribes/value/"
)

// BribesKey returns the store key to retrieve a Bribes from the index fields
func BribesKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
