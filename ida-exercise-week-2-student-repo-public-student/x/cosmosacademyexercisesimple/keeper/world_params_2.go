package keeper

import (
	"encoding/binary"

	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetWorldParams2Count get the total number of worldParams2
func (k Keeper) GetWorldParams2Count(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WorldParams2CountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWorldParams2Count set the total number of worldParams2
func (k Keeper) SetWorldParams2Count(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WorldParams2CountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWorldParams2 appends a worldParams2 in the store with a new id and update the count
func (k Keeper) AppendWorldParams2(
	ctx sdk.Context,
	worldParams2 types.WorldParams2,
) uint64 {
	// Create the worldParams2
	count := k.GetWorldParams2Count(ctx)

	// Set the ID of the appended value
	worldParams2.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParams2Key))
	appendedValue := k.cdc.MustMarshal(&worldParams2)
	store.Set(GetWorldParams2IDBytes(worldParams2.Id), appendedValue)

	// Update worldParams2 count
	k.SetWorldParams2Count(ctx, count+1)

	return count
}

// SetWorldParams2 set a specific worldParams2 in the store
func (k Keeper) SetWorldParams2(ctx sdk.Context, worldParams2 types.WorldParams2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParams2Key))
	b := k.cdc.MustMarshal(&worldParams2)
	store.Set(GetWorldParams2IDBytes(worldParams2.Id), b)
}

// GetWorldParams2 returns a worldParams2 from its id
func (k Keeper) GetWorldParams2(ctx sdk.Context, id uint64) (val types.WorldParams2, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParams2Key))
	b := store.Get(GetWorldParams2IDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWorldParams2 removes a worldParams2 from the store
func (k Keeper) RemoveWorldParams2(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParams2Key))
	store.Delete(GetWorldParams2IDBytes(id))
}

// GetAllWorldParams2 returns all worldParams2
func (k Keeper) GetAllWorldParams2(ctx sdk.Context) (list []types.WorldParams2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParams2Key))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WorldParams2
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetWorldParams2IDBytes returns the byte representation of the ID
func GetWorldParams2IDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetWorldParams2IDFromBytes returns ID in uint64 format from a byte array
func GetWorldParams2IDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
