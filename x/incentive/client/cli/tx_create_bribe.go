package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/ibc-go/v6/modules/core/04-channel/client/utils"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"hhand/x/incentive/types"
)

var _ = strconv.Itoa(0)

func CmdSendCreateBribe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-create-bribe [src-port] [src-channel] [proposer] [title] [block] [denom] [amount] [chain]",
		Short: "Send a create-bribe over IBC",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argProposer := args[2]
			argTitle := args[3]
			argBlock, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argDenom := args[5]
			argAmount, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}
			argChain, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendCreateBribe(creator, srcPort, srcChannel, timeoutTimestamp, argProposer, argTitle, argBlock, argDenom, argAmount, argChain)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
