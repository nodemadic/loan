package loan_test

import (
	"testing"

	keepertest "github.com/nodemadic/loan/testutil/keeper"
	"github.com/nodemadic/loan/testutil/nullify"
	"github.com/nodemadic/loan/x/loan"
	"github.com/nodemadic/loan/x/loan/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoanKeeper(t)
	loan.InitGenesis(ctx, *k, genesisState)
	got := loan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
