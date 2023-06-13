package gaws

import (
	"fmt"

	"github.com/future-jim/gaws/pkg/gaws"
	"github.com/spf13/cobra"
)

var deleteObjectsCmd = &cobra.Command{
	Use:     "delete-objects",
	Aliases: []string{"del-obj"},
	Short:   "permanently deletes all objects in an s3 bucket",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		deleteObjects()

	},
}

func init() {
	rootCmd.AddCommand(deleteObjectsCmd)
}

func deleteObjects() {

	gaws.ListBuckets()

	bucketPromptContentName := promptContent{
		"Please enter a bucket name.",
		"Enter the name of the bucket you would like to empty of ALL objects",
	}

	bucket := promptGetInput(bucketPromptContentName)
	fmt.Printf("The bucket name you entered is %s\n\n", bucket)

	bucketPromptContentConfirm := promptContent{
		"Yes or No",
		"Confirm deletion of all objects (y/n)",
	}

	confirmation := promptGetInput(bucketPromptContentConfirm)

	if confirmation == "y" {
		gaws.DeleteObjects(bucket)
	} else {
		println("delete-objects not confirmed and exited")
	}
}
