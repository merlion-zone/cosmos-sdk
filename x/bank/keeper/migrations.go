package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// v043 "github.com/cosmos/cosmos-sdk/x/bank/legacy/v043"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	// return v043.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
	// NOTE: bypass migration from version 1 to 2, since we always start from version 2.
	return nil
}
