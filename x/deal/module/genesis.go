package deal

import (
	"example/x/deal/keeper"
	"example/x/deal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.DealCounter != nil {
		k.SetDealCounter(ctx, *genState.DealCounter)
	}
	// Set all the newDeal
	for _, elem := range genState.NewDealList {
		k.SetNewDeal(ctx, elem)
	}

	for _, elem := range genState.ContractCounter {
		k.SetContractCounter(ctx, *elem)

	}
	// Set all the newContract
	for _, elem := range genState.NewContractList {
		k.SetNewContract(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all dealCounter
	dealCounter, found := k.GetDealCounter(ctx)
	if found {
		genesis.DealCounter = &dealCounter
	}
	genesis.NewDealList = k.GetAllNewDeal(ctx)
	// Get all contractCounter
	contractCounter, err := k.GetAllContractCounter(ctx)
	if err == nil {
		genesis.ContractCounter = contractCounter
	}

	for _, counter := range contractCounter {
		contractsForDealId := k.GetAllNewContract(ctx, counter.DealId)
		genesis.NewContractList = append(genesis.NewContractList, contractsForDealId...)
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
