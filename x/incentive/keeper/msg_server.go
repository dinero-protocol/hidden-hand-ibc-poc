package keeper

import (
	"context"
	"hhand/x/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) DistributeBribe(c context.Context, msg *types.MsgDistributeBribeRequest) (*types.MsgDistributeBribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	bribes := m.GetAllBribes(ctx)
	totalCoins := sdk.NewCoins()
	for _, bribe := range bribes {
		totalCoins = totalCoins.Add(sdk.NewCoin(bribe.Denom, sdk.NewInt(int64(bribe.Amount))))
		// Delete the bribe.
		m.RemoveBribes(ctx, bribe.Index)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"DistributeBribes",
			sdk.NewAttribute("total_bribes", totalCoins.String()),
		),
	)

	return &types.MsgDistributeBribeResponse{}, nil
}
