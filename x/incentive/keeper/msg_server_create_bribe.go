package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"hhand/x/incentive/types"
)

func (k msgServer) SendCreateBribe(goCtx context.Context, msg *types.MsgSendCreateBribe) (*types.MsgSendCreateBribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.CreateBribePacketData

	packet.Proposer = msg.Proposer
	packet.Title = msg.Title
	packet.Block = msg.Block
	packet.Denom = msg.Denom
	packet.Amount = msg.Amount
	packet.Chain = msg.Chain

	// Transmit the packet
	_, err := k.TransmitCreateBribePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendCreateBribeResponse{}, nil
}
