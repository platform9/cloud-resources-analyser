package cmd

import (
	"fmt"
	"os"

	colour "github.com/platform9/cloud-analyser/pkg/color"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "cloud-analyser",
	Long:             `CLI for cloud resources analyser i.e AWS`,
	PersistentPreRun: ensureSecrets,
}

// To make API secrets are loaded.
func ensureSecrets(cmd *cobra.Command, args []string) {
	if cmd.Name() == "help" || cmd.Name() == "version" {
		return
	}
	requiredSecrets := map[string]string{
		"AWS_REGION":            os.Getenv("AWS_REGION"),
		"AWS_SECRET_ACCESS_KEY": os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"AWS_ACCESS_KEY_ID":     os.Getenv("AWS_ACCESS_KEY_ID"),
	}
	missing := ""
	for secret, val := range requiredSecrets {
		if val == "" {
			missing = fmt.Sprintf("%s%s ", missing, secret)
		}
	}
	if missing != "" {
		fmt.Printf("%s\n", colour.Red("cloud-analyser secrets not set. Please ensure the following values are set:\n",
			missing,
		))
		os.Exit(1)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		zap.S().Fatalf(err.Error())
	}
}

func init() {
	// To tell Cobra not to provide the default completion command.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
