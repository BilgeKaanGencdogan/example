package keeper

import (
	"example/x/deal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Then use it in the keeper like this:
func (k Keeper) SetNewDeal(ctx sdk.Context, newDeal types.NewDeal) {
	store := k.storeService.OpenKVStore(ctx)

	b := k.cdc.MustMarshal(&newDeal)
	if err := store.Set(types.NewDealStoreKey(newDeal.DealId), b); err != nil {
		panic(err) // or handle error appropriately
	}
}

func (k Keeper) GetNewDeal(
	ctx sdk.Context,
	index string,
) (val types.NewDeal, found bool) {
	store := k.storeService.OpenKVStore(ctx)

	// Create the full key by combining prefix and index
	key := append(types.KeyPrefix(types.NewDealKeyPrefix), types.NewDealKey(index)...)

	b, err := store.Get(key)
	if err != nil {
		return val, false
	}

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewDeal removes a newDeal from the store
func (k Keeper) RemoveNewDeal(
	ctx sdk.Context,
	index string,

) {
	store := k.storeService.OpenKVStore(ctx)
	// Create the full key by combining prefix and index
	key := append(types.KeyPrefix(types.NewDealKeyPrefix), types.NewDealKey(index)...)

	store.Delete(key)
}

func (k Keeper) GetAllNewDeal(ctx sdk.Context) (list []types.NewDeal) {
	store := k.storeService.OpenKVStore(ctx)

	// Get iterator using the NewDealKeyPrefix
	iterator, err := store.Iterator(types.KeyPrefix(types.NewDealKeyPrefix), nil)
	if err != nil {
		return nil
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NewDeal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
