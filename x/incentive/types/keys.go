package types

import fmt "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "incentive"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_incentive"

	// Version defines the current version the IBC module supports
	Version = "incentive-1"

	// PortID is the default port id that module binds to
	PortID = "incentive"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("incentive-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// Create the bribe index key.
func BribeIndex(portID string, channelID string, proposer string) string {
	return fmt.Sprintf("%s-%s-%s", portID, channelID, proposer)
}
