package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "example/testutil/keeper"
	"example/x/deal/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DealKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
