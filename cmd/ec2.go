package cmd

import (
	"github.com/platform9/cloud-analyser/pkg/logic"
	"github.com/spf13/cobra"
)

var listEc2 = `
  # Get all EC2 instances in a given region.
  Ex: cloud-analyser listEc2
  `
var listEc2i = `
  # Get  EC2 instances in a given region.
  Ex: cloud-analyser listEc2i
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
	all bool
	//region string
)

func init() {
	rootCmd.AddCommand(listEc2Cmd)
	//appCmdPodRes.Flags().StringVarP(&region, "region", "r", "", "AWS Region")
	listEc2Cmd.Flags().BoolVarP(&all, "all", "a", false, "To fetch all EC2 instances")
}

func listEc2Run(cmd *cobra.Command, args []string) {

	if all {
		logic.DescribeInstancesCmd()
	} else {
		logic.GetInstancebykeyCmd()
	}
}
