package gaws

import (
	
	"os"
	"fmt"
	"errors"
	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
	"github.com/future-jim/gaws/pkg/gaws"

)

var diruploadCmd = &cobra.Command{
    Use:   "dup",
    Aliases: []string{"directory-upload"},
    Short:  "dup",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
        dirupload()
        
    },
}

func init() {
    rootCmd.AddCommand(diruploadCmd)
}

type promptContent struct {
	errorMsg string
	label    string
}

func dirupload() {
	bucketPromptContent := promptContent{
		"Please enter a bucket name",
		"What bucket would you like to upload to?",
	}

	buckets, err := gaws.BucketList()
	if err != nil {
		fmt.Printf("Couldnt list buckets: %v", err)
	}
	println("\nOr enter a new bucket name:\n")

	bucket := promptGetInput(bucketPromptContent)
	fmt.Printf("The bucket name you entered is %s\n\n", bucket)


	for _, bucket := range buckets.Buckets{
		fmt.Printf("%s\n", *bucket.Name)
		//:TODO CHECK IF BUCKET ALREADY EXISTS, IF EXISTS DONT TRY TO CREATE IT
		}
	
		
	
	directoryPromptContent := promptContent{
		"Please enter a local directory",
		"What local directory would you like to upload?",
	}
	directory := promptGetInput(directoryPromptContent)
	fmt.Printf("The directory you selected is %s\n\n", directory)

	gaws.Tarfunc()
	
	gaws.CreateBucket(bucket)

	
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	//	fmt.Printf("Input: %s\n", result)

	return result
}
