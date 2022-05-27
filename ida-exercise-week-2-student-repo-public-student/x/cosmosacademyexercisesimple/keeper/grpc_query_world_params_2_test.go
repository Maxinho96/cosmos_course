package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/b9lab/cosmos-academy-exercise-simple/testutil/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
)

func TestWorldParams2QuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWorldParams2(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetWorldParams2Request
		response *types.QueryGetWorldParams2Response
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetWorldParams2Request{Id: msgs[0].Id},
			response: &types.QueryGetWorldParams2Response{WorldParams2: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetWorldParams2Request{Id: msgs[1].Id},
			response: &types.QueryGetWorldParams2Response{WorldParams2: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetWorldParams2Request{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.WorldParams2(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestWorldParams2QueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWorldParams2(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWorldParams2Request {
		return &types.QueryAllWorldParams2Request{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WorldParams2All(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WorldParams2), step)
			require.Subset(t, msgs, resp.WorldParams2)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WorldParams2All(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WorldParams2), step)
			require.Subset(t, msgs, resp.WorldParams2)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WorldParams2All(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WorldParams2All(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
