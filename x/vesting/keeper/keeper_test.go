package keeper_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
<<<<<<< HEAD
	"github.com/evmos/evmos/v19/app"
	"github.com/evmos/evmos/v19/encoding"
	"github.com/evmos/evmos/v19/testutil/integration/evmos/network"
	testutiltx "github.com/evmos/evmos/v19/testutil/tx"
	"github.com/evmos/evmos/v19/x/vesting/keeper"
	vestingtypes "github.com/evmos/evmos/v19/x/vesting/types"
	"github.com/stretchr/testify/require"
=======
	"github.com/evmos/evmos/v19/app"
	"github.com/evmos/evmos/v19/encoding"
	"github.com/evmos/evmos/v19/x/vesting/keeper"
	vestingtypes "github.com/evmos/evmos/v19/x/vesting/types"
>>>>>>> main
)

func TestNewKeeper(t *testing.T) {
	nw := network.NewUnitTestNetwork()
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	cdc := encCfg.Codec
	storeKey := storetypes.NewKVStoreKey(vestingtypes.StoreKey)

	addr, _ := testutiltx.NewAccAddressAndKey()

	testcases := []struct {
		name      string
		authority sdk.AccAddress
		expPass   bool
	}{
		{
			name:      "valid authority format",
			authority: addr,
			expPass:   true,
		},
		{
			name:      "empty authority",
			authority: []byte{},
			expPass:   false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expPass {
				newKeeper := keeper.NewKeeper(
					storeKey,
					tc.authority,
					cdc,
<<<<<<< HEAD
					nw.App.AccountKeeper,
					nw.App.BankKeeper,
					nw.App.DistrKeeper,
					nw.App.EvmKeeper,
					nw.App.StakingKeeper,
					nw.App.GovKeeper,
=======
					suite.app.AccountKeeper,
					suite.app.BankKeeper,
					suite.app.DistrKeeper,
					suite.app.EvmKeeper,
					suite.app.StakingKeeper,
					suite.app.GovKeeper,
>>>>>>> main
				)
				require.NotNil(t, newKeeper)
			} else {
				require.PanicsWithError(t, "addresses cannot be empty: unknown address", func() {
					_ = keeper.NewKeeper(
						storeKey,
						tc.authority,
						cdc,
<<<<<<< HEAD
						nw.App.AccountKeeper,
						nw.App.BankKeeper,
						nw.App.DistrKeeper,
						nw.App.EvmKeeper,
						nw.App.StakingKeeper,
						nw.App.GovKeeper,
=======
						suite.app.AccountKeeper,
						suite.app.BankKeeper,
						suite.app.DistrKeeper,
						suite.app.EvmKeeper,
						suite.app.StakingKeeper,
						suite.app.GovKeeper,
>>>>>>> main
					)
				})
			}
		})
	}
}
