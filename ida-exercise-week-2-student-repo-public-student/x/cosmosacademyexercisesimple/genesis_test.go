package cosmosacademyexercisesimple_test

import (
	"testing"

	keepertest "github.com/b9lab/cosmos-academy-exercise-simple/testutil/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/testutil/nullify"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WorldParams2List: []types.WorldParams2{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		WorldParams2Count: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	cosmosacademyexercisesimple.InitGenesis(ctx, *k, genesisState)
	got := cosmosacademyexercisesimple.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Len(t, got.WorldParams2List, len(genesisState.WorldParams2List))
	require.Subset(t, genesisState.WorldParams2List, got.WorldParams2List)
	require.Equal(t, genesisState.WorldParams2Count, got.WorldParams2Count)
	// this line is used by starport scaffolding # genesis/test/assert
}
