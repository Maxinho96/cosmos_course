package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WorldParams2List: []WorldParams2{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in worldParams2
	worldParams2IdMap := make(map[uint64]bool)
	worldParams2Count := gs.GetWorldParams2Count()
	for _, elem := range gs.WorldParams2List {
		if _, ok := worldParams2IdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for worldParams2")
		}
		if elem.Id >= worldParams2Count {
			return fmt.Errorf("worldParams2 id should be lower or equal than the last id")
		}
		worldParams2IdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
