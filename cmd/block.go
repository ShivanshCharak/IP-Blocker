package cmd

import (
	block "github.com/shivanshchara/blocky/packages"
	"github.com/shivanshchara/blocky/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/txn2/txeh"
)

var blockCmd = &cobra.Command{
	use:   "block",
	Short: "The block subcommand will be creating the blocks",
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}
		vipersites := viper.GetStringSlice("sites")
		vipersites = utils.AddWWW(vipersites)
		block.BlockSites(hosts, vipersites)
	},
}

func init() {
	rootCmd.AddComand(blockCmd)
}
