package cosmosacademyexercisesimple

import (
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the worldParams2
	for _, elem := range genState.WorldParams2List {
		k.SetWorldParams2(ctx, elem)
	}

	// Set worldParams2 count
	k.SetWorldParams2Count(ctx, genState.WorldParams2Count)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WorldParams2List = k.GetAllWorldParams2(ctx)
	genesis.WorldParams2Count = k.GetWorldParams2Count(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
