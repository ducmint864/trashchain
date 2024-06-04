package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "trashchain/testutil/keeper"
	"trashchain/x/trashchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TrashchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
