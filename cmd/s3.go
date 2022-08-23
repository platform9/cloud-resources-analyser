package cmd

import (
	"github.com/platform9/cloud-analyser/pkg/logic"
	"github.com/spf13/cobra"
)

var list_s3 = `
  # Get all s3 instances in a given region.
  Ex: cloud-analyser lists3
  `

var (
	lists3Cmd = &cobra.Command{
		Use:     "lists3",
		Short:   "Show all the s3 buckets in a given region",
		Example: list_s3,
		Long:    `"Show all the running s3 buckets in a given region"`,
		Run:     lists3Run,
	}
)

func init() {
	rootCmd.AddCommand(lists3Cmd)
	//appCmdPodRes.Flags().StringVarP(&region, "region", "r", "", "AWS Region")
	//listEc2Cmd.Flags().BoolVarP(&all, "all", "a", true, "To fetch all EC2 instances")
}

func lists3Run(cmd *cobra.Command, args []string) {

	logic.AccountBucketOps()

}
