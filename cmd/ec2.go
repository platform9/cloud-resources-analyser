package cmd

import (
	"github.com/platform9/cloud-analyser/pkg/logic"
	"github.com/spf13/cobra"
)

var list_ec2 = `
  # Get all EC2 instances in a given region.
  Ex: cloud-analyser listEc2
  `

var (
	listEc2Cmd = &cobra.Command{
		Use:     "listEc2",
		Short:   "Show all the running EC2 instances in a given region",
		Example: list_ec2,
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
	listEc2Cmd.Flags().BoolVarP(&all, "all", "a", true, "To fetch all EC2 instances")
}

func listEc2Run(cmd *cobra.Command, args []string) {

	if all {
		logic.DescribeInstancesCmd()
	}
}
