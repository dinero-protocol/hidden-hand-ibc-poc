package keeper

import (
	"errors"

	"hhand/x/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// TransmitCreateBribePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitCreateBribePacket(
	ctx sdk.Context,
	packetData types.CreateBribePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvCreateBribePacket processes packet reception
func (k Keeper) OnRecvCreateBribePacket(ctx sdk.Context, packet channeltypes.Packet, data types.CreateBribePacketData) (packetAck types.CreateBribePacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// Check if the bribe already exists.
	index := types.BribeIndex(packet.GetSourcePort(), packet.GetSourceChannel(), data.Proposer)
	_, found := k.GetBribes(ctx, index)
	if found {
		return packetAck, errors.New("bribe already exists")
	}

	// Check if the coin that came in is from this chain.
	denom, isSaved := k.OriginalDenom(ctx, packet.DestinationPort, packet.DestinationChannel, data.Denom)
	if !isSaved {
		denom = VoucherDenom(packet.SourcePort, packet.SourceChannel, data.Denom)
	}

	// Mint the tokens to the module account.
	if err := k.bankKeeper.MintCoins(
		ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(int64(data.Amount)))),
	); err != nil {
		return packetAck, err
	}

	bribe := types.Bribes{
		Proposer: data.Proposer,
		Title:    data.Title,
		Block:    data.Block,
		Denom:    data.Denom,
		Amount:   data.Amount,
		Chain:    data.Chain,
	}
	bribe.Index = index

	// Set the bribe
	k.SetBribes(ctx, bribe)

	return packetAck, nil
}

// OnAcknowledgementCreateBribePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementCreateBribePacket(ctx sdk.Context, packet channeltypes.Packet, data types.CreateBribePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.CreateBribePacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// Set the Bribe.
		index := types.BribeIndex(packet.GetSourcePort(), packet.GetSourceChannel(), data.Proposer)
		bribe := types.Bribes{
			Proposer: data.Proposer,
			Title:    data.Title,
			Block:    data.Block,
			Denom:    data.Denom,
			Amount:   data.Amount,
			Chain:    data.Chain,
		}
		bribe.Index = index
		k.SetBribes(ctx, bribe)

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutCreateBribePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutCreateBribePacket(ctx sdk.Context, packet channeltypes.Packet, data types.CreateBribePacketData) error {

	// TODO: packet timeout logic

	return nil
}
