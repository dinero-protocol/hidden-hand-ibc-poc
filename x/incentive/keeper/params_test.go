package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "hhand/testutil/keeper"
	"hhand/x/incentive/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IncentiveKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
