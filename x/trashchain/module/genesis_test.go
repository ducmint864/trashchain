package trashchain_test

import (
	"testing"

	keepertest "trashchain/testutil/keeper"
	"trashchain/testutil/nullify"
	trashchain "trashchain/x/trashchain/module"
	"trashchain/x/trashchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TrashchainKeeper(t)
	trashchain.InitGenesis(ctx, k, genesisState)
	got := trashchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
