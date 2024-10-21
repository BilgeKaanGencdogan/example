package keeper

import (
	"example/x/deal/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetContractCounter set contractCounter in the store
func (k Keeper) SetContractCounter(ctx sdk.Context, contractCounter types.ContractCounter) {
	store := k.storeService.OpenKVStore(ctx)
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))
	b := k.cdc.MustMarshal(&contractCounter)
	key := createKey([]byte(types.ContractCounterKey), []byte(contractCounter.DealId))
	store.Set(key, b)
}

// GetContractCounter returns contractCounter
func (k Keeper) GetContractCounter(ctx sdk.Context, dealId string) (val types.ContractCounter, found bool) {
	store := k.storeService.OpenKVStore(ctx)
	key := createKey(types.KeyPrefix(types.ContractCounterKey), []byte(dealId))
	b, err := store.Get(key)
	if err != nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// GetAllContractCounter gets all the contract counter from store
func (k Keeper) GetAllContractCounter(ctx sdk.Context) ([]*types.ContractCounter, error) {
	store := k.storeService.OpenKVStore(ctx)

	// Initialize the slice with capacity
	contractCounters := make([]*types.ContractCounter, 0)

	// Create an iterator using store's Iterator method
	iterator, err := store.Iterator(types.KeyPrefix(types.ContractCounterKey), nil) // nil as prefix end
	if err != nil {
		return nil, fmt.Errorf("failed to create iterator: %w", err)
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var counter types.ContractCounter
		value := iterator.Value()

		if err := k.cdc.Unmarshal(value, &counter); err != nil {
			return nil, fmt.Errorf("failed to unmarshal contract counter: %w", err)
		}

		contractCounters = append(contractCounters, &counter)
	}

	return contractCounters, nil
}

// RemoveContractCounter removes contractCounter from the store
func (k Keeper) RemoveContractCounter(ctx sdk.Context, dealId string) {
	store := k.storeService.OpenKVStore(ctx)
	key := createKey(types.KeyPrefix(types.ContractCounterKey), []byte(dealId))
	store.Delete(key)
}
