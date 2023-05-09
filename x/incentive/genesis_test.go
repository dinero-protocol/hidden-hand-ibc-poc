package incentive_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hhand/testutil/keeper"
	"hhand/testutil/nullify"
	"hhand/x/incentive"
	"hhand/x/incentive/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		BribesList: []types.Bribes{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		DenomTraceList: []types.DenomTrace{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IncentiveKeeper(t)
	incentive.InitGenesis(ctx, *k, genesisState)
	got := incentive.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.BribesList, got.BribesList)
	require.ElementsMatch(t, genesisState.DenomTraceList, got.DenomTraceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
