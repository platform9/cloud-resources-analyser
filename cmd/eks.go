package cmd

import (
	"github.com/platform9/cloud-analyser/pkg/logic"
	"github.com/spf13/cobra"
)

var listEks = `
  # Get all EC2 instances in a given region.
  Ex: cloud-analyser listEc2
  `

var (
	listEksCmd = &cobra.Command{
		Use:     "listEks",
		Short:   "Display running eks clusters",
		Example: listEks,
		Long:    `"Show all the running eks clusters in a given region"`,
		Run:     listEksRun,
	}
)

func init() {
	rootCmd.AddCommand(listEksCmd)
}

func listEksRun(cmd *cobra.Command, args []string) {

	logic.ListEKSClusters()
}
