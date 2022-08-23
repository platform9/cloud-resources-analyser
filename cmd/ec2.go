package cmd

import (
	"github.com/platform9/cloud-analyser/pkg/logic"
	"github.com/spf13/cobra"
)

var listEc2 = `
  # Get all EC2 instances in a given region.
  Ex: cloud-analyser listEc2

  # Get EC2 instances by key
  Ex: cloud-analyser listEc2 -k <KEY>
`

var (
	listEc2Cmd = &cobra.Command{
		Use:     "listEc2",
		Short:   "Show all the running EC2 instances in a given region",
		Example: listEc2,
		Long:    `"Show all the running EC2 instances in a given region"`,
		Run:     listEc2Run,
	}
)

var (
	keyName string
)

func init() {
	rootCmd.AddCommand(listEc2Cmd)
	listEc2Cmd.Flags().StringVarP(&keyName, "key", "k", "", "Key Name to filter")
}

func listEc2Run(cmd *cobra.Command, args []string) {

	if keyName == "" {
		logic.DescribeInstancesCmd("")
	} else {
		logic.DescribeInstancesCmd(keyName)
	}
}
