package cli

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-errors/errors"

	"github.com/spf13/cobra"
)

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
	newOAIClients()

}
