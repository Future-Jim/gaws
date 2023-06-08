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
    Use:   "directory-upload",
    Aliases: []string{"directory-upload"},
    Short:  "uploads a directory as a tar archive to a specified S3 bucket",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
        archiveUpload()
        
    },
}

func init() {
    rootCmd.AddCommand(diruploadCmd)
}

type promptContent struct {
	errorMsg string
	label    string
}

func archiveUpload() {
	bucketPromptContent := promptContent{
		"Please enter a bucket name.",
		"Enter bucket to upload archive to",
	}

	buckets, err := gaws.BucketList()
	if err != nil {
		fmt.Printf("Couldnt list buckets: %v", err)
	}
	println("\nOr enter a new bucket name:\n")

	bucket := promptGetInput(bucketPromptContent)
	fmt.Printf("The bucket name you entered is %s\n\n", bucket)

	directoryPromptContent := promptContent{
		"Please enter a local directory",
		"Enter the relative path of directory to archive and upload",
	}
	
	directory := promptGetInput(directoryPromptContent)
	fmt.Printf("The directory you selected is %s\n\n", directory)
	
	filenamePromptContent := promptContent{
		"Please enter a filename for the archive",
		"Enter filename for archive:",
	}

	filename := promptGetInput(filenamePromptContent)
	fmt.Printf("The directory you selected is %s\n\n", directory)
	err = os.Remove(filename)
	if err != nil {
		fmt.Println(err)
	}
	
	file := gaws.CreateTar(filename, directory)

	for _, buckets := range buckets.Buckets {
		   if bucket == *buckets.Name {
			gaws.S3Fileupload(filename, file, bucket)			
			return
		}		
	}
	gaws.CreateBucket(bucket)
	gaws.S3Fileupload(filename, file, bucket)
	return 
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

	return result
}
