package keeper

import (
	"trashchain/x/trashchain/types"
)

var _ types.QueryServer = Keeper{}
