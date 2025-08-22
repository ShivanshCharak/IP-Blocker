package cmd

import (
	"fmt"

	block "github.com/shivanshchara/blocky/packages"
	"github.com/shivanshchara/blocky/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/txn2/txeh"
)

var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Unblock given sites",
	Long:  "Unblock provided sites",
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}
		viperUnblockSites := viper.GetStringMapStringSlice("sites")
		viperUnblockSites = utils.AddWWW(viperUnblockSites)
		fmt.Println(viperUnblockSites)
		block.CleanBlocks(hosts, viperUnblockSites)
	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)
}
