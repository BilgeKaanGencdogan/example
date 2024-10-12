package keeper

import (
	"example/x/deal/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNewContract set a specific newContract in the store -  "NewContract/value/{dealId}"
func (k Keeper) SetNewContract(ctx sdk.Context, newContract types.NewContract) {
	key := types.NewContractKey(newContract.DealId)        // {byte} key of "NewContract/value/{dealId}/"
	store := prefix.NewStore(ctx.KVStore(k.storeKey), key) // {this returns store object} "store{parent: KVStore, prefix: NewContract/value/{dealId}/}"
	b := k.cdc.MustMarshal(&newContract)                   // The serialized byte array of object
	store.Set(types.NewContractKey(
		newContract.ContractId,
	), b) // {this sets the value in the store} "NewContract/value/{dealId}/{contractId}"
}

// GetNewContract returns a newContract from its index
func (k Keeper) GetNewContract(
	ctx sdk.Context,
	dealId string,
	contractId string,
) (val types.NewContract, found bool) {
	storeKey := types.NewContractKey(dealId)                    // {byte} key of "NewContract/value/{dealId}/"
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey) //{this returns store object} // "store{parent: KVStore, prefix: NewContract/value/{dealId}/}"
	b := store.Get(types.NewContractKey(contractId))
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
	storeKey := types.NewContractKey(dealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)
	store.Delete(types.NewContractKey(contractId))
}

// GetAllNewContract returns all newContract
func (k Keeper) GetAllNewContract(ctx sdk.Context, dealId string) (list []types.NewContract) {
	storeKey := types.NewContractKey(dealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NewContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
