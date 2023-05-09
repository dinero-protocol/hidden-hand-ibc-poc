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

	// Mock distribution.
	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	bribes := m.GetAllBribes(ctx)
	totalCoins := sdk.NewCoins()
	for _, bribe := range bribes {
		totalCoins = totalCoins.Add(sdk.NewCoin(bribe.Denom, sdk.NewInt(int64(bribe.Amount))))
		// Delete the bribe.
		m.RemoveBribes(ctx, bribe.Index)
	}

	if err := m.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, totalCoins); err != nil {
		return nil, err
	}

	return &types.MsgDistributeBribeResponse{}, nil
}
