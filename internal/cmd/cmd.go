package cmd

import (
	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

var wf *aw.Workflow

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	wf = aw.New()
	rootCmd.AddCommand(decodeCmd)
	rootCmd.AddCommand(encodeCmd)
	rootCmd.AddCommand(signCmd)
	rootCmd.AddCommand(dateCmd)
	rootCmd.AddCommand(uniqueIdCmd)
	rootCmd.AddCommand(netCmd)
}
