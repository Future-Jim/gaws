package gaws

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:  "gaws",
    Short: "go + aws",
    Long: "gaws - a simple CLI that helps with some AWS tasks",
  

    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}
