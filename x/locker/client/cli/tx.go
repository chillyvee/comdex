package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/comdex-official/comdex/x/locker/types"
)

func txCreateLocker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-locker [app_mapping_id] [asset_id] [amount]",
		Short: "create a new locker",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			appMappingId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			assetId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return types.ErrorInvalidAmountIn
			}

			msg := types.NewMsgCreateLockerRequest(ctx.FromAddress, amount, assetId, appMappingId)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txDepositAssetLocker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-locker [locker_id] [amount] [asset_id] [app_mapping_id] ",
		Short: "deposit to a locker",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			lockerId := args[0]
			if len(lockerId) == 0 {
				return err
			}
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return types.ErrorInvalidAmountIn
			}
			assetId, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			appMappingId, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositAssetRequest(ctx.FromAddress, lockerId, amount, assetId, appMappingId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txWithdrawAssetLocker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-locker [locker_id] [amount] [asset_id] [app_mapping_id] ",
		Short: "dwithdraw from a locker",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			lockerId := args[0]
			if len(lockerId) == 0 {
				return err
			}
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return types.ErrorInvalidAmountIn
			}
			assetId, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			appMappingId, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawAssetRequest(ctx.FromAddress, lockerId, amount, assetId, appMappingId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txAddWhiteListedAssetLocker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist-asset-locker [app_mapping_id][asset_id] ",
		Short: "withdraw from a locker",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			appMappingId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			assetId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddWhiteListedAssetRequest(ctx.FromAddress, appMappingId, assetId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}