package keeper_test

import (
<<<<<<< HEAD
	testkeyring "github.com/evmos/evmos/v18/testutil/integration/evmos/keyring"
	"github.com/evmos/evmos/v18/testutil/integration/evmos/network"
	evmtypes "github.com/evmos/evmos/v18/x/evm/types"
=======
	"github.com/cometbft/cometbft/abci/types"
	evmtypes "github.com/evmos/evmos/v19/x/evm/types"
>>>>>>> main
)

func (suite *KeeperTestSuite) TestEndBlock() {
	keyring := testkeyring.New(2)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	ctx := unitNetwork.GetContext()
	preEventManager := ctx.EventManager()
	suite.Require().Equal(0, len(preEventManager.Events()))

	err := unitNetwork.App.EvmKeeper.EndBlock(ctx)
	suite.Require().NoError(err)

	postEventManager := unitNetwork.GetContext().EventManager()
	// should emit 1 EventTypeBlockBloom event on EndBlock
	suite.Require().Equal(1, len(postEventManager.Events()))
	suite.Require().Equal(evmtypes.EventTypeBlockBloom, postEventManager.Events()[0].Type)
}
