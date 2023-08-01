// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity >=0.8.18;

import "../common/Types.sol";

/// @dev The VestingI contract's address.
address constant VESTING_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000803;

/// @dev The VestingI contract's instance.
VestingI constant VESTING_CONTRACT = VestingI(VESTING_PRECOMPILE_ADDRESS);

/// @dev Define all the available staking methods.
string constant MSG_CREATE_CLAWBACK_VESTING_ACCOUNT = "/evmos.vesting.v2.MsgCreateClawbackVestingAccount";
string constant MSG_FUND_VESTING_ACCOUNT = "/evmos.vesting.v2.MsgFundVestingAccount";
string constant MSG_CLAWBACK = "/evmos.vesting.v2.MsgClawback";
string constant MSG_CONVERT_VESTING_ACCOUNT = "/evmos.vesting.v2.MsgConvertVestingAccount";
string constant MSG_UPDATE_VESTING_FUNDER = "/evmos.vesting.v2.MsgUpdateVestingFunder";


// Period defines a length of time and amount of coins that will vest.
struct Period {
    int64 length;
    Coin[] amount;
}

/// @author Evmos Team
/// @title Vesting Precompiled Contract
/// @dev The interface through which solidity contracts will interact with vesting.
/// We follow this same interface including four-byte function selectors, in the precompile that
/// wraps the pallet.
/// @custom:address 0x0000000000000000000000000000000000000803
interface VestingI {
    /// @dev Defines a method for creating a new clawback vesting account.
    /// @param funderAddress The address of the account that will fund the vesting account.
    /// @param vestingAddress The address of the account that will receive the vesting account.
    function createClawbackVestingAccount(
        address funderAddress,
        address vestingAddress
    ) external returns (bool success);

    /// @dev Defines a method for funding a vesting account.
    /// @param funderAddress The address of the account that will fund the vesting account.
    /// @param vestingAddress The address of the clawback vesting account that will receive the vesting funds.
    /// @param startTime The time at which the vesting account will start.
    /// @param lockupPeriods The lockup periods of the vesting account.
    /// @param vestingPeriods The vesting periods of the vesting account.
    function fundVestingAccount(
        address funderAddress,
        address vestingAddress,
        uint64 startTime,
        Period[] calldata lockupPeriods,
        Period[] calldata vestingPeriods
    ) external returns (bool success);

    /// @dev Defines a method for clawing back coins from a vesting account.
    /// @param funderAddress The address of the account that funded the vesting account.
    /// @param accountAddress The address of the vesting account.
    /// @param destAddress The address of the account that will receive the clawed back coins.
    function clawback(
        address funderAddress,
        address accountAddress,
        address destAddress
    ) external returns (bool success);

    /// @dev Defines a method for updating the funder of a vesting account.
    /// @param funderAddress The address of the account that funded the vesting account.
    /// @param newFunderAddress The address of the new funder of the vesting account.
    /// @param vestingAddress The address of the vesting account.
    function updateVestingFunder(
        address funderAddress,
        address newFunderAddress,
        address vestingAddress
    ) external returns (bool success);

    /// @dev Defines a method for converting a vesting account to a clawback vesting account.
    /// @param vestingAddress The address of the vesting account.
    function convertVestingAccount(
        address vestingAddress
    ) external returns (bool success);

    /// QUERIES

    /// @dev Defines a query for getting the balances of a vesting account.
    /// @param vestingAddress The address of the vesting account.
    function balances(
        address vestingAddress
    ) external view returns (Coin[] memory locked, Coin[] memory unvested, Coin[] memory vested);

    /// @dev Defines an event that is emitted when a clawback vesting account is created.
    /// @param funderAddress The address of the account that funded the vesting account.
    /// @param vestingAddress The address of the account that received the vesting account.
    event CreateClawbackVestingAccount(
        address indexed funderAddress,
        address indexed vestingAddress
    );

    /// @dev Defines an event that is emitted when a clawback vesting account is funded.
    /// @param funderAddress The address of the account that funded the vesting account.
    /// @param vestingAddress The address of the account that received the vesting account.
    /// @param startTime The time at which the vesting account will start.
    /// @param lockupPeriods The lockup periods of the vesting account.
    /// @param vestingPeriods The vesting periods of the vesting account.
    event FundVestingAccount(
        address indexed funderAddress,
        address indexed vestingAddress,
        uint64 startTime,
        Period[] lockupPeriods,
        Period[] vestingPeriods
    );

    /// @dev Defines an event that is emitted when a vesting account is clawed back.
    /// @param funderAddress The address of the account that funded the vesting account.
    /// @param accountAddress The address of the vesting account.
    /// @param destAddress The address of the account that received the clawed back coins.
    event Clawback(
        address indexed funderAddress,
        address indexed accountAddress,
        address destAddress
    );

    /// @dev Defines an event that is emitted when a vesting account's funder is updated.
    /// @param funderAddress The address of the account that funded the vesting account.
    /// @param newFunderAddress The address of the new funder of the vesting account.
    /// @param vestingAddress The address of the vesting account.
    event UpdateVestingFunder(
        address indexed funderAddress,
        address indexed vestingAddress,
        address  newFunderAddress
    );

    /// @dev Defines an event that is emitted when a vesting account is converted to a clawback vesting account.
    /// @param vestingAddress The address of the vesting account.
    event ConvertVestingAccount(
        address indexed vestingAddress
    );
}