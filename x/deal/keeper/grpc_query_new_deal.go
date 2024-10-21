package keeper

import (
	"context"
	"example/x/deal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewDealAll(c context.Context, req *types.QueryAllNewDealRequest) (*types.QueryAllNewDealResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newDeals []types.NewDeal
	ctx := sdk.UnwrapSDKContext(c)
	store := k.storeService.OpenKVStore(ctx)

	// Get iterator with NewDealKeyPrefix
	iterator, err := store.Iterator(types.KeyPrefix(types.NewDealKeyPrefix), nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer iterator.Close()

	// Handle pagination
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
		if len(newDeals) >= limit {
			break
		}

		var newDeal types.NewDeal
		if err := k.cdc.Unmarshal(iterator.Value(), &newDeal); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		newDeals = append(newDeals, newDeal)
		count++
	}

	// Create pagination response
	pageRes := &query.PageResponse{
		NextKey: nil,
		Total:   uint64(count),
	}

	return &types.QueryAllNewDealResponse{
		NewDeal:    newDeals,
		Pagination: pageRes,
	}, nil
}

// NewDeal is the query handler to fetch the deal details for a given dealId
func (k Keeper) NewDeal(c context.Context, req *types.QueryGetNewDealRequest) (*types.QueryGetNewDealResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNewDeal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetNewDealResponse{NewDeal: val}, nil
}
