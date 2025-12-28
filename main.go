package main

import (
	"log"

	"github.com/RajaPremSai/terraform-ai-go/cmd/cli"
	"github.com/RajaPremSai/terraform-ai-go/pkg/utils"
)

func main() {
	workingDir, err := utils.CurrentDir()
	if err != nil {
		log.Fatalf("Unable to fetch current dir,%s\n", err.Error())
	}

	execDir, err := utils.TerraformPath()
	if err != nil {
		log.Fatalf("Unable to execute dir:%s\n", err.Error())
	}
	cli.InitAndExecute(workingDir, execDir)
}
