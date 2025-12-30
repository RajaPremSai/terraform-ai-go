package cli

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/RajaPremSai/terraform-ai-go/pkg/terraform"
	"github.com/RajaPremSai/terraform-ai-go/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const initSubCommand = "You are a Terraform HCL generator, only generate valid provider Terraform HCL templates."

var errLength = errors.New("invalid length")

func addInit() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Run terraform init",
		RunE:  initCommand,
	}
	return initCmd
}

func initCommand(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.Wrap(errLength, "prompt must be provided")
	}
	return initCmd(args)
}

func initCmd(args []string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	oaiClients, err := newOAIClients()
	if err != nil {
		return fmt.Errorf("error creating new OAI Client:%w", err)
	}
	var action, com string
	for action != apply {
		args = append(args, action)
		com, err = completion(ctx, oaiClients, args, *openAIDeploymentName, initSubCommand)
		if err != nil {
			return fmt.Errorf("error completion:%w", err)
		}
		text := fmt.Sprintf("/n Attempting to apply the following template:%s", com)
		log.Println(text)

		action, err = userActionPrompt()
		if err != nil {
			return err
		}
		if action == dontApply {
			return nil
		}
	}
	if err = terraform.CheckTemplate(com); err != nil {
		return fmt.Errorf("error checking template:%w", err)
	}

	if err = utils.StoreFile("provider.tf", com); err != nil {
		return fmt.Errorf("error storing file:%w", err)
	}
	if err = ops.Init(); err != nil {
		return fmt.Errorf("error running terraform init:%w", err)
	}

	return nil
}
