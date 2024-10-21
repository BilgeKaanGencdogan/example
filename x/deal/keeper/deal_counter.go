package keeper

import (
	"example/x/deal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDealCounter set dealCounter in the store
func (k Keeper) SetDealCounter(ctx sdk.Context, dealCounter types.DealCounter) {
	store := k.storeService.OpenKVStore(ctx)
	b := k.cdc.MustMarshal(&dealCounter)
	store.Set([]byte{0}, b)
}

// GetDealCounter returns dealCounter
func (k Keeper) GetDealCounter(ctx sdk.Context) (val types.DealCounter, found bool) {
	store := k.storeService.OpenKVStore(ctx)

	// Retrieve the value from the store
	b, err := store.Get([]byte{0})
	if err != nil {
		// Handle the error if store.Get fails
		return val, false
	}

	// Check if the value is nil
	if b == nil {
		return val, false
	}

	// Unmarshal the value into the DealCounter struct
	err = k.cdc.Unmarshal(b, &val)
	if err != nil {
		// Handle the error if unmarshalling fails
		return val, false
	}

	return val, true
}

func (k Keeper) RemoveDealCounter(ctx sdk.Context) {
	store := k.storeService.OpenKVStore(ctx)

	store.Delete([]byte{0})
}
