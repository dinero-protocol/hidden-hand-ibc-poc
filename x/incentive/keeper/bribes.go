package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hhand/x/incentive/types"
)

// SetBribes set a specific bribes in the store from its index
func (k Keeper) SetBribes(ctx sdk.Context, bribes types.Bribes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BribesKeyPrefix))
	b := k.cdc.MustMarshal(&bribes)
	store.Set(types.BribesKey(
		bribes.Index,
	), b)
}

// GetBribes returns a bribes from its index
func (k Keeper) GetBribes(
	ctx sdk.Context,
	index string,

) (val types.Bribes, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BribesKeyPrefix))

	b := store.Get(types.BribesKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBribes removes a bribes from the store
func (k Keeper) RemoveBribes(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BribesKeyPrefix))
	store.Delete(types.BribesKey(
		index,
	))
}

// GetAllBribes returns all bribes
func (k Keeper) GetAllBribes(ctx sdk.Context) (list []types.Bribes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BribesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bribes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
