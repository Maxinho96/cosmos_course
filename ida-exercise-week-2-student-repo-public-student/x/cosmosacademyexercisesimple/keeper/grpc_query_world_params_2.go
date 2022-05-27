package keeper

import (
	"context"

	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WorldParams2All(c context.Context, req *types.QueryAllWorldParams2Request) (*types.QueryAllWorldParams2Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var worldParams2s []types.WorldParams2
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	worldParams2Store := prefix.NewStore(store, types.KeyPrefix(types.WorldParams2Key))

	pageRes, err := query.Paginate(worldParams2Store, req.Pagination, func(key []byte, value []byte) error {
		var worldParams2 types.WorldParams2
		if err := k.cdc.Unmarshal(value, &worldParams2); err != nil {
			return err
		}

		worldParams2s = append(worldParams2s, worldParams2)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWorldParams2Response{WorldParams2: worldParams2s, Pagination: pageRes}, nil
}

func (k Keeper) WorldParams2(c context.Context, req *types.QueryGetWorldParams2Request) (*types.QueryGetWorldParams2Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	worldParams2, found := k.GetWorldParams2(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetWorldParams2Response{WorldParams2: worldParams2}, nil
}
