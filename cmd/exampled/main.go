package main

import (
	"fmt"
	"os"

	clienthelpers "cosmossdk.io/client/v2/helpers"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"example/app"
	"example/cmd/exampled/cmd"
)

func main() {
	// Create a new codec
	appCodec := codec.NewProtoCodec(types.NewInterfaceRegistry())

	rootCmd := cmd.NewRootCmd(appCodec)
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
