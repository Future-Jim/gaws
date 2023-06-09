package gaws

import (
	"github.com/future-jim/gaws/pkg/gaws"
	"github.com/spf13/cobra"
)

var listBucketsCmd = &cobra.Command{
	Use:     "list-buckets",
	Aliases: []string{"list-buckets"},
	Short:   "list all buckets from s3",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		gaws.ListBuckets()

	},
}

func init() {
	rootCmd.AddCommand(listBucketsCmd)

}
