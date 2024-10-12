package keeper

import (
	"example/x/deal/types"
)

var _ types.QueryServer = Keeper{}
