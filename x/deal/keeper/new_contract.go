package keeper

import (
	"example/x/deal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNewContract set a specific newContract in the store -  "NewContract/value/{dealId}"
func (k Keeper) SetNewContract(ctx sdk.Context, newContract types.NewContract) {
	// Create a hierarchical key structure manually
	dealPrefix := types.NewContractKey(newContract.DealId)
	// Ensure proper separation between dealId and contractId in the key
	fullKey := append(dealPrefix, []byte("/"+newContract.ContractId)...)

	store := k.storeService.OpenKVStore(ctx)
	b := k.cdc.MustMarshal(&newContract) // The serialized byte array of object
	store.Set(fullKey, b)                // {this sets the value in the store} "NewContract/value/{dealId}/{contractId}"
}

// GetNewContract returns a newContract from its index
func (k Keeper) GetNewContract(
	ctx sdk.Context,
	dealId string,
	contractId string,
) (val types.NewContract, found bool) {
	// Create a hierarchical key structure manually
	dealPrefix := types.NewContractKey(dealId)
	// Ensure proper separation between dealId and contractId in the key
	fullKey := append(dealPrefix, []byte("/"+contractId)...)

	store := k.storeService.OpenKVStore(ctx)

	b, err := store.Get(fullKey)
	if err != nil {
		return val, false
	}

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveNewContract(
	ctx sdk.Context,
	dealId string,
	contractId string,
) {
	// Create a hierarchical key structure manually
	dealPrefix := types.NewContractKey(dealId)
	// Ensure proper separation between dealId and contractId in the key
	fullKey := append(dealPrefix, []byte("/"+contractId)...)

	store := k.storeService.OpenKVStore(ctx)

	store.Delete(fullKey)
}

func (k Keeper) GetAllNewContract(ctx sdk.Context, dealId string) (list []types.NewContract) {
	// Create the prefix key for the deal
	dealPrefix := types.NewContractKey(dealId)

	store := k.storeService.OpenKVStore(ctx)

	// Use the store's iterator with the deal prefix
	iterator, err := store.Iterator(dealPrefix, nil)
	if err != nil {
		return nil
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NewContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
