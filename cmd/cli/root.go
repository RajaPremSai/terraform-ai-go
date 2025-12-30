package cli

import (
	"flag"
	"log"
	"strconv"

	terraform "github.com/RajaPremSai/terraform-ai-go/pkg/terraform"
	"github.com/spf13/cobra"
	"github.com/walles/env"
)

const (
	verion = "0.0.2"
)

var (
	openAIDeploymentName = flag.String("openai-deplyment-name", env.GetOr("OPENAI_DEPLOYMENT_NAME", env.String, "text-davinci-003"), "The deployment name used for the model in OpenAI Service")
	openAIPIKey          = flag.String("openai-api-key", env.GetOr("OPENAI_API_KEY", env.String, ""), "It is API Key from Open AI - REQUIRED")
	workingDir           = flag.String("working-dir", env.GetOr("WORKING_DIR", env.String, ""), "The path where the project we want to run")
	execDir              = flag.String("exe-dir", env.GetOr("EXEC_DIR", env.String, ""), "The path of terraform")
	requireConfirmation  = flag.Bool("required-confirmation", env.GetOr("REQUIRED_CONFIRMATION", strconv.ParseBool, true), "whether to reuire confirmation before executing the command.Defaults to true")
	azureOpenAIEndpoint  = flag.String("azure-openai-endpoint", env.GetOr("AZURE_OPENAI_ENDPOINT", env.String, ""), "The endpoint for azure openai service.If provided, Azure OpenAI service will be used instead of OpenAI service.")
	ops                  terraform.Ops
	err                  error
	temperature          = flag.Float64("temperature", env.GetOr("TEMPERATURE", env.WithBitSize(strconv.ParseFloat, 64), 0.0), "The temperature to use for the model.")
)

func InitAndExecute(workDir string, executionDir string) {
	flag.Parse()

	if *workingDir == "" {
		workingDir = &workDir
	}

	if *execDir == "" {
		execDir = &executionDir
	}

	if *openAIPIKey == "" {
		log.Fatal("Please provide Open AI API Key ")
	}

	if err := RootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

func RootCmd() *cobra.Command {
	ops, err = terraform.NewTerraform(*workingDir, *execDir)
	if err != nil {
		return nil
	}
	cmd := &cobra.Command{
		Use:          "terraform-assistant",
		Version:      verion,
		Args:         cobra.MinimumNArgs(1),
		RunE:         runCommand,
		SilenceUsage: true,
	}

	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	initCmd := addInit()
	cmd.AddCommand(initCmd)

	return cmd
}
