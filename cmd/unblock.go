package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	log "github.com/visionmedia/go-cli-log"

	"github.com/tdepaulo/bacenjud/internal/client"
)

var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "unblock the user's account related with process",
	Run:   ExecuteUnblock,
	Args:  cobra.ExactArgs(2),
}

func ExecuteUnblock(cmd *cobra.Command, args []string) {
	var protocolId = args[0]
	var authToken = args[1]

	log.Width = "0"
	log.Verbose = true

	log.Info("searching", "retrieving the process %s", protocolId)

	// 1. First, we get the blockIds by Protocol
	blockIds, err := client.GetBlockIds(authToken, protocolId)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// 2. Filter the blockIds that is already unblocked
	for _, blockId := range blockIds {
		log.Info("validating", "validate if %s is already unblocked", blockId)

		unblocked, err := client.IsUnblocked(authToken, blockId)
		if err != nil {
			log.Error(err)
			continue
		}

		if unblocked {
			log.Warn(fmt.Sprintf("the unblock of %s is already done", blockId))
			continue
		}

		// 3. And do the unblock itself
		log.Info("unblocking", "blockId %s", blockId)
		if err := client.Unblock(authToken, blockId); err != nil {
			log.Error(err)
			continue
		}
	}
}
