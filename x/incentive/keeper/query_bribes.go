package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hhand/x/incentive/types"
)

func (k Keeper) BribesAll(goCtx context.Context, req *types.QueryAllBribesRequest) (*types.QueryAllBribesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bribess []types.Bribes
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	bribesStore := prefix.NewStore(store, types.KeyPrefix(types.BribesKeyPrefix))

	pageRes, err := query.Paginate(bribesStore, req.Pagination, func(key []byte, value []byte) error {
		var bribes types.Bribes
		if err := k.cdc.Unmarshal(value, &bribes); err != nil {
			return err
		}

		bribess = append(bribess, bribes)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBribesResponse{Bribes: bribess, Pagination: pageRes}, nil
}

func (k Keeper) Bribes(goCtx context.Context, req *types.QueryGetBribesRequest) (*types.QueryGetBribesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetBribes(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBribesResponse{Bribes: val}, nil
}
