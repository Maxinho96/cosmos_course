package keeper_test

import (
	"testing"

	keepertest "github.com/b9lab/cosmos-academy-exercise-simple/testutil/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNWorldParams2(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WorldParams2 {
	items := make([]types.WorldParams2, n)
	for i := range items {
		items[i].Id = keeper.AppendWorldParams2(ctx, items[i])
	}
	return items
}

func TestWorldParams2Get(t *testing.T) {
	keeper, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	items := createNWorldParams2(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetWorldParams2(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestWorldParams2Remove(t *testing.T) {
	keeper, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	items := createNWorldParams2(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWorldParams2(ctx, item.Id)
		_, found := keeper.GetWorldParams2(ctx, item.Id)
		require.False(t, found)
	}
}

func TestWorldParams2GetAll(t *testing.T) {
	keeper, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	items := createNWorldParams2(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllWorldParams2(ctx))
}

func TestWorldParams2Count(t *testing.T) {
	keeper, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	items := createNWorldParams2(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetWorldParams2Count(ctx))
}
