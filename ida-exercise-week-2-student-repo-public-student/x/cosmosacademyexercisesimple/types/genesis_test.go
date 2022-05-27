package types_test

import (
	"testing"

	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				WorldParams2List: []types.WorldParams2{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				WorldParams2Count: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated worldParams2",
			genState: &types.GenesisState{
				WorldParams2List: []types.WorldParams2{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid worldParams2 count",
			genState: &types.GenesisState{
				WorldParams2List: []types.WorldParams2{
					{
						Id: 1,
					},
				},
				WorldParams2Count: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
