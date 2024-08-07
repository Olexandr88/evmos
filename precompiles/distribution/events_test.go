package distribution_test

import (
	"math/big"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
<<<<<<< HEAD
	"github.com/evmos/evmos/v18/cmd/config"
	cmn "github.com/evmos/evmos/v18/precompiles/common"
	"github.com/evmos/evmos/v18/precompiles/distribution"
	"github.com/evmos/evmos/v18/utils"
	"github.com/evmos/evmos/v18/x/evm/core/vm"
	"github.com/evmos/evmos/v18/x/evm/statedb"
=======
	"github.com/evmos/evmos/v19/cmd/config"
	cmn "github.com/evmos/evmos/v19/precompiles/common"
	"github.com/evmos/evmos/v19/precompiles/distribution"
	"github.com/evmos/evmos/v19/utils"
	"github.com/evmos/evmos/v19/x/evm/core/vm"
>>>>>>> main
)

func (s *PrecompileTestSuite) TestSetWithdrawAddressEvent() {
	var (
		ctx  sdk.Context
		stDB *statedb.StateDB
	)
	method := s.precompile.Methods[distribution.SetWithdrawAddressMethod]
	testCases := []struct {
		name        string
		malleate    func(operatorAddress string) []interface{}
		postCheck   func()
		gas         uint64
		expError    bool
		errContains string
	}{
		{
			"success - the correct event is emitted",
			func(string) []interface{} {
				return []interface{}{
					s.keyring.GetAddr(0),
					s.keyring.GetAddr(0).String(),
				}
			},
			func() {
				log := stDB.Logs()[0]
				s.Require().Equal(log.Address, s.precompile.Address())

				// Check event signature matches the one emitted
				event := s.precompile.ABI.Events[distribution.EventTypeSetWithdrawAddress]
				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight()))

				// Check the fully unpacked event matches the one emitted
				var setWithdrawerAddrEvent distribution.EventSetWithdrawAddress
				err := cmn.UnpackLog(s.precompile.ABI, &setWithdrawerAddrEvent, distribution.EventTypeSetWithdrawAddress, *log)
				s.Require().NoError(err)
				s.Require().Equal(s.keyring.GetAddr(0), setWithdrawerAddrEvent.Caller)
				s.Require().Equal(sdk.MustBech32ifyAddressBytes(config.Bech32Prefix, s.keyring.GetAddr(0).Bytes()), setWithdrawerAddrEvent.WithdrawerAddress)
			},
			20000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.SetupTest()
		ctx = s.network.GetContext()
		stDB = s.network.GetStateDB()

		contract := vm.NewContract(vm.AccountRef(s.keyring.GetAddr(0)), s.precompile, big.NewInt(0), tc.gas)
		ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		initialGas := ctx.GasMeter().GasConsumed()
		s.Require().Zero(initialGas)

		_, err := s.precompile.SetWithdrawAddress(ctx, s.keyring.GetAddr(0), contract, stDB, &method, tc.malleate(s.network.GetValidators()[0].OperatorAddress))

		if tc.expError {
			s.Require().Error(err)
			s.Require().Contains(err.Error(), tc.errContains)
		} else {
			s.Require().NoError(err)
			tc.postCheck()
		}
	}
}

func (s *PrecompileTestSuite) TestWithdrawDelegatorRewardsEvent() {
	var (
		ctx  sdk.Context
		stDB *statedb.StateDB
	)
	method := s.precompile.Methods[distribution.WithdrawDelegatorRewardsMethod]
	testCases := []struct {
		name        string
		malleate    func(val stakingtypes.Validator) []interface{}
		postCheck   func()
		gas         uint64
		expError    bool
		errContains string
	}{
		{
			"success - the correct event is emitted",
			func(val stakingtypes.Validator) []interface{} {
				var err error

				ctx, err = s.prepareStakingRewards(ctx, stakingRewards{
					Validator: val,
					Delegator: s.keyring.GetAccAddr(0),
					RewardAmt: testRewardsAmt,
				})
				s.Require().NoError(err)
				return []interface{}{
					s.keyring.GetAddr(0),
					val.OperatorAddress,
				}
			},
			func() {
				log := stDB.Logs()[0]
				s.Require().Equal(log.Address, s.precompile.Address())

				// Check event signature matches the one emitted
				event := s.precompile.ABI.Events[distribution.EventTypeWithdrawDelegatorRewards]
				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight()))

				optAddr, err := sdk.ValAddressFromBech32(s.network.GetValidators()[0].OperatorAddress)
				s.Require().NoError(err)
				optHexAddr := common.BytesToAddress(optAddr)

				// Check the fully unpacked event matches the one emitted
				var delegatorRewards distribution.EventWithdrawDelegatorRewards
				err = cmn.UnpackLog(s.precompile.ABI, &delegatorRewards, distribution.EventTypeWithdrawDelegatorRewards, *log)
				s.Require().NoError(err)
				s.Require().Equal(s.keyring.GetAddr(0), delegatorRewards.DelegatorAddress)
				s.Require().Equal(optHexAddr, delegatorRewards.ValidatorAddress)
				s.Require().Equal(expRewardsAmt.BigInt(), delegatorRewards.Amount)
			},
			20000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.SetupTest()
		ctx = s.network.GetContext()
		stDB = s.network.GetStateDB()

		contract := vm.NewContract(vm.AccountRef(s.keyring.GetAddr(0)), s.precompile, big.NewInt(0), tc.gas)
		ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		initialGas := ctx.GasMeter().GasConsumed()
		s.Require().Zero(initialGas)

		_, err := s.precompile.WithdrawDelegatorRewards(ctx, s.keyring.GetAddr(0), contract, stDB, &method, tc.malleate(s.network.GetValidators()[0]))

		if tc.expError {
			s.Require().Error(err)
			s.Require().Contains(err.Error(), tc.errContains)
		} else {
			s.Require().NoError(err)
			tc.postCheck()
		}
	}
}

func (s *PrecompileTestSuite) TestWithdrawValidatorCommissionEvent() {
	var (
		ctx  sdk.Context
		stDB *statedb.StateDB
		amt  = math.NewInt(1e18)
	)
	method := s.precompile.Methods[distribution.WithdrawValidatorCommissionMethod]
	testCases := []struct {
		name        string
		malleate    func(operatorAddress string) []interface{}
		postCheck   func()
		gas         uint64
		expError    bool
		errContains string
	}{
		{
			"success - the correct event is emitted",
			func(operatorAddress string) []interface{} {
				valAddr, err := sdk.ValAddressFromBech32(operatorAddress)
				s.Require().NoError(err)
				valCommission := sdk.DecCoins{sdk.NewDecCoinFromDec(utils.BaseDenom, math.LegacyNewDecFromInt(amt))}
				// set outstanding rewards
				s.Require().NoError(s.network.App.DistrKeeper.SetValidatorOutstandingRewards(ctx, valAddr, types.ValidatorOutstandingRewards{Rewards: valCommission}))
				// set commission
				s.Require().NoError(s.network.App.DistrKeeper.SetValidatorAccumulatedCommission(ctx, valAddr, types.ValidatorAccumulatedCommission{Commission: valCommission}))
				// set funds to distr mod to pay for commission
				coins := sdk.NewCoins(sdk.NewCoin(utils.BaseDenom, amt))
				err = s.mintCoinsForDistrMod(ctx, coins)
				s.Require().NoError(err)
				return []interface{}{
					operatorAddress,
				}
			},
			func() {
				log := stDB.Logs()[0]
				s.Require().Equal(log.Address, s.precompile.Address())

				// Check event signature matches the one emitted
				event := s.precompile.ABI.Events[distribution.EventTypeWithdrawValidatorCommission]
				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight()))

				// Check the fully unpacked event matches the one emitted
				var validatorRewards distribution.EventWithdrawValidatorRewards
				err := cmn.UnpackLog(s.precompile.ABI, &validatorRewards, distribution.EventTypeWithdrawValidatorCommission, *log)
				s.Require().NoError(err)
				s.Require().Equal(crypto.Keccak256Hash([]byte(s.network.GetValidators()[0].OperatorAddress)), validatorRewards.ValidatorAddress)
				s.Require().Equal(amt.BigInt(), validatorRewards.Commission)
			},
			20000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.SetupTest()
		ctx = s.network.GetContext()
		stDB = s.network.GetStateDB()

		valAddr, err := sdk.ValAddressFromBech32(s.network.GetValidators()[0].GetOperator())
		s.Require().NoError(err)
		validatorAddress := common.BytesToAddress(valAddr)
		contract := vm.NewContract(vm.AccountRef(validatorAddress), s.precompile, big.NewInt(0), tc.gas)
		ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		initialGas := ctx.GasMeter().GasConsumed()
		s.Require().Zero(initialGas)

		_, err = s.precompile.WithdrawValidatorCommission(ctx, validatorAddress, contract, stDB, &method, tc.malleate(s.network.GetValidators()[0].OperatorAddress))

		if tc.expError {
			s.Require().Error(err)
			s.Require().Contains(err.Error(), tc.errContains)
		} else {
			s.Require().NoError(err)
			tc.postCheck()
		}
	}
}

//nolint:dupl
func (s *PrecompileTestSuite) TestClaimRewardsEvent() {
	var (
		ctx  sdk.Context
		stDB *statedb.StateDB
	)
	testCases := []struct {
		name      string
		coins     sdk.Coins
		postCheck func()
	}{
		{
			"success",
			sdk.NewCoins(sdk.NewCoin(utils.BaseDenom, math.NewInt(1e18))),
			func() {
				log := stDB.Logs()[0]
				s.Require().Equal(log.Address, s.precompile.Address())
				// Check event signature matches the one emitted
				event := s.precompile.ABI.Events[distribution.EventTypeClaimRewards]
				s.Require().Equal(event.ID, common.HexToHash(log.Topics[0].Hex()))
				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight()))

				var claimRewardsEvent distribution.EventClaimRewards
				err := cmn.UnpackLog(s.precompile.ABI, &claimRewardsEvent, distribution.EventTypeClaimRewards, *log)
				s.Require().NoError(err)
				s.Require().Equal(common.BytesToAddress(s.keyring.GetAddr(0).Bytes()), claimRewardsEvent.DelegatorAddress)
				s.Require().Equal(big.NewInt(1e18), claimRewardsEvent.Amount)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()
			ctx = s.network.GetContext()
			stDB = s.network.GetStateDB()
			err := s.precompile.EmitClaimRewardsEvent(ctx, stDB, s.keyring.GetAddr(0), tc.coins)
			s.Require().NoError(err)
			tc.postCheck()
		})
	}
}

func (s *PrecompileTestSuite) TestFundCommunityPoolEvent() {
	var (
		ctx  sdk.Context
		stDB *statedb.StateDB
	)
	testCases := []struct {
		name      string
		coins     sdk.Coins
		postCheck func()
	}{
		{
			"success - the correct event is emitted",
			sdk.NewCoins(sdk.NewCoin(utils.BaseDenom, math.NewInt(1e18))),
			func() {
				log := stDB.Logs()[0]
				s.Require().Equal(log.Address, s.precompile.Address())
				// Check event signature matches the one emitted
				event := s.precompile.ABI.Events[distribution.EventTypeFundCommunityPool]
				s.Require().Equal(event.ID, common.HexToHash(log.Topics[0].Hex()))
				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight()))

				var fundCommunityPoolEvent distribution.EventFundCommunityPool
				err := cmn.UnpackLog(s.precompile.ABI, &fundCommunityPoolEvent, distribution.EventTypeFundCommunityPool, *log)
				s.Require().NoError(err)
				s.Require().Equal(s.keyring.GetAddr(0), fundCommunityPoolEvent.Depositor)
				s.Require().Equal(big.NewInt(1e18), fundCommunityPoolEvent.Amount)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()
			ctx = s.network.GetContext()
			stDB = s.network.GetStateDB()

			err := s.precompile.EmitFundCommunityPoolEvent(ctx, stDB, s.keyring.GetAddr(0), tc.coins)
			s.Require().NoError(err)
			tc.postCheck()
		})
	}
}

//nolint:dupl
func (s *PrecompileTestSuite) TestFundCommunityPoolEvent() {
	testCases := []struct {
		name      string
		coins     sdk.Coins
		postCheck func()
	}{
		{
			"success - the correct event is emitted",
			sdk.NewCoins(sdk.NewCoin(utils.BaseDenom, math.NewInt(1e18))),
			func() {
				log := s.stateDB.Logs()[0]
				s.Require().Equal(log.Address, s.precompile.Address())
				// Check event signature matches the one emitted
				event := s.precompile.ABI.Events[distribution.EventTypeFundCommunityPool]
				s.Require().Equal(event.ID, common.HexToHash(log.Topics[0].Hex()))
				s.Require().Equal(log.BlockNumber, uint64(s.ctx.BlockHeight()))

				var fundCommunityPoolEvent distribution.EventFundCommunityPool
				err := cmn.UnpackLog(s.precompile.ABI, &fundCommunityPoolEvent, distribution.EventTypeFundCommunityPool, *log)
				s.Require().NoError(err)
				s.Require().Equal(common.BytesToAddress(s.address.Bytes()), fundCommunityPoolEvent.Depositor)
				s.Require().Equal(big.NewInt(1e18), fundCommunityPoolEvent.Amount)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()

			err := s.precompile.EmitFundCommunityPoolEvent(s.ctx, s.stateDB, s.address, tc.coins)
			s.Require().NoError(err)
			tc.postCheck()
		})
	}
}
