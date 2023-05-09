package types

// ValidateBasic is used for validating the packet
func (p CreateBribePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p CreateBribePacketData) GetBytes() ([]byte, error) {
	var modulePacket IncentivePacketData

	modulePacket.Packet = &IncentivePacketData_CreateBribePacket{&p}

	return modulePacket.Marshal()
}
