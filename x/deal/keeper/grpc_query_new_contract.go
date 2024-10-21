package keeper

import (
	"context"
	"example/x/deal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewContractAll(c context.Context, req *types.QueryAllNewContractRequest) (*types.QueryAllNewContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newContracts []types.NewContract
	ctx := sdk.UnwrapSDKContext(c)
	store := k.storeService.OpenKVStore(ctx)

	// Create the prefix for contracts under this dealId
	prefixKey := types.NewContractKey(req.DealId)

	// Get iterator with prefix
	iterator, err := store.Iterator(prefixKey, nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer iterator.Close()

	// Handle pagination manually
	start := 0
	if req.Pagination != nil && req.Pagination.Offset > 0 {
		start = int(req.Pagination.Offset)
	}

	limit := 100 // default limit
	if req.Pagination != nil && req.Pagination.Limit > 0 {
		limit = int(req.Pagination.Limit)
	}

	count := 0
	for ; iterator.Valid(); iterator.Next() {
		// Skip entries before start
		if count < start {
			count++
			continue
		}
		// Break if we've reached the limit
		if len(newContracts) >= limit {
			break
		}

		var newContract types.NewContract
		if err := k.cdc.Unmarshal(iterator.Value(), &newContract); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		newContracts = append(newContracts, newContract)
		count++
	}

	// Create pagination response
	pageRes := &query.PageResponse{
		NextKey: nil, // Set to nil since we're using offset pagination
		Total:   uint64(count),
	}

	return &types.QueryAllNewContractResponse{
		NewContract: newContracts,
		Pagination:  pageRes,
	}, nil
}

// NewContract is the query handler to fatch the contract for a given specific dealId and contractId
func (k Keeper) NewContract(c context.Context, req *types.QueryGetNewContractRequest) (*types.QueryGetNewContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNewContract(
		ctx,
		req.DealId,
		req.ContractId,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
	return &types.QueryGetNewContractResponse{NewContract: val}, nil
}
