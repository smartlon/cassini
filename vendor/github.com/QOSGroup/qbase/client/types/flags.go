package types

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagChainID   = "chain-id"
	FlagNode      = "node"
	FlagHeight    = "height"
	FlagNonce     = "nonce"
	FlagAsync     = "async"
	FlagTrustNode = "trust-node"
	FlagMaxGas    = "max-gas"
	FlagJSONIndet = "indent"
	FlagNonceNode = "nonce-node"

	//qcp flag
	FlagQcp            = "qcp" //启用QCP模式,发送txQcp消息
	FlagQcpSigner      = "qcp-signer"
	FlagQcpSequence    = "qcp-seq"
	FlagQcpFrom        = "qcp-from"
	FlagQcpBlockHeight = "qcp-blockheight"
	FlagQcpTxIndex     = "qcp-txindex"
	FlagQcpExtends     = "qcp-extends"
)

// LineBreak can be included in a command list to provide a blank line
// to help with readability
var (
	LineBreak = &cobra.Command{Run: func(*cobra.Command, []string) {}}
)

// GetCommands adds common flags to query commands
func GetCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		c.Flags().String(FlagChainID, "", "Chain ID of tendermint node")
		c.Flags().Bool(FlagTrustNode, false, "Trust connected full node (don't verify proofs for responses)")
		c.Flags().String(FlagNode, "tcp://localhost:26657", "<host>:<port> to tendermint rpc interface for this chain")
		c.Flags().Int64(FlagHeight, 0, "block height to query, omit to get most recent provable block")
		c.Flags().Bool(FlagJSONIndet, false, "add indent to json response")
		viper.BindPFlag(FlagChainID, c.Flags().Lookup(FlagChainID))
		viper.BindPFlag(FlagNode, c.Flags().Lookup(FlagNode))
	}
	return cmds
}

// PostCommands adds common flags for commands to post tx
func PostCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		c.Flags().Int64(FlagNonce, 0, "account nonce to sign the tx")
		c.Flags().Int64(FlagMaxGas, 0, "gas limit to set per tx")
		c.Flags().String(FlagChainID, "", "Chain ID of tendermint node")
		c.Flags().String(FlagNode, "tcp://localhost:26657", "<host>:<port> to tendermint rpc interface for this chain")
		c.Flags().Bool(FlagAsync, false, "broadcast transactions asynchronously")
		c.Flags().Bool(FlagTrustNode, false, "Trust connected full node (don't verify proofs for responses)")
		c.Flags().Bool(FlagQcp, false, "enable qcp mode. send qcp tx")
		c.Flags().String(FlagQcpSigner, "", "qcp mode flag. qcp tx signer key name")
		c.Flags().String(FlagQcpFrom, "", "qcp mode flag. qcp tx source chainID")
		c.Flags().Int64(FlagQcpSequence, 0, "qcp mode flag.  qcp in sequence")
		c.Flags().Int64(FlagQcpBlockHeight, 0, "qcp mode flag. original tx blockheight, blockheight must greater than 0")
		c.Flags().Int64(FlagQcpTxIndex, 0, "qcp mode flag. original tx index")
		c.Flags().String(FlagQcpExtends, "", "qcp mode flag. qcp tx extends info")
		c.Flags().Bool(FlagJSONIndet, false, "add indent to json response")
		c.Flags().String(FlagNonceNode, "", "tcp://<host>:<port> to tendermint rpc interface for some chain to query account nonce")

		viper.BindPFlag(FlagChainID, c.Flags().Lookup(FlagChainID))
		viper.BindPFlag(FlagMaxGas, c.Flags().Lookup(FlagMaxGas))
		viper.BindPFlag(FlagNode, c.Flags().Lookup(FlagNode))
	}
	return cmds
}
