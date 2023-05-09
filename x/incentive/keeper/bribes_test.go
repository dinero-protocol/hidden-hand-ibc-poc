package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "hhand/testutil/keeper"
	"hhand/testutil/nullify"
	"hhand/x/incentive/keeper"
	"hhand/x/incentive/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBribes(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Bribes {
	items := make([]types.Bribes, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetBribes(ctx, items[i])
	}
	return items
}

func TestBribesGet(t *testing.T) {
	keeper, ctx := keepertest.IncentiveKeeper(t)
	items := createNBribes(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBribes(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBribesRemove(t *testing.T) {
	keeper, ctx := keepertest.IncentiveKeeper(t)
	items := createNBribes(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBribes(ctx,
			item.Index,
		)
		_, found := keeper.GetBribes(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestBribesGetAll(t *testing.T) {
	keeper, ctx := keepertest.IncentiveKeeper(t)
	items := createNBribes(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBribes(ctx)),
	)
}
