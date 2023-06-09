package gaws

import (
	"fmt"

	"github.com/future-jim/gaws/pkg/gaws"
	"github.com/spf13/cobra"
)

var deleteBucketCmd = &cobra.Command{
	Use:     "delete-bucket",
	Aliases: []string{"del-bucket"},
	Short:   "permanently deletes all an s3 bucket",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		deleteBucket()

	},
}

func init() {
	rootCmd.AddCommand(deleteBucketCmd)
}

func deleteBucket() {

	gaws.ListBuckets()
	bucketPromptContentName := promptContent{
		"Please enter a bucket name.",
		"Enter the name of the bucket you would like to delete",
	}

	bucket := promptGetInput(bucketPromptContentName)
	fmt.Printf("\nThe bucket name you entered is %s\n\n", bucket)

	bucketPromptContentConfirm := promptContent{
		"Yes or No",
		"Confirm deletion of the bucket (y/n)",
	}

	confirmation := promptGetInput(bucketPromptContentConfirm)

	if confirmation == "y" {
		gaws.DeleteBucket(bucket)
	} else {
		println("delete-bucket not confirmed and exited")
	}
}
