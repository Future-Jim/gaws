package gaws

import (
	"github.com/future-jim/gaws/pkg/gaws"
	"github.com/spf13/cobra"
)

var bucketListCmd = &cobra.Command{
	Use:     "bucketlist",
	Aliases: []string{"bl"},
	Short:   "list buckets from s3",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		gaws.BucketList()

	},
}

func init() {
	rootCmd.AddCommand(bucketListCmd)

}
