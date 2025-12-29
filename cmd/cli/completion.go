package cli

import (
	"context"
	"fmt"
	"regexp"

	openai "github.com/PullRequestInc/go-gpt3"
	azureopenai "github.com/RajaPremSai/terraform-ai-go/pkg/gpt3"
	"github.com/pkg/errors"
	gptEncoder "github.com/samber/go-gpt-3-encoder"
)

type oaiClients struct {
	azureClient  azureopenai.Client
	openAIClient openai.Client
}

const userRole = "user"

func newOAIClients() (oaiClients, error) {
	var (
		oaiClient   openai.Client
		azureClient azureopenai.Client
		err         error
	)
	if azureOpenAIEndpoint == nil || *azureOpenAIEndpoint == "" {
		oaiClient = openai.NewClient(*openAIPIKey)
	} else {
		re := regexp.MustCompile(`^[a-zA-Z0-9]+([_-]?[a-zA-Z0-9]+)*$`)
		if !re.MatchString(*openAIDeploymentName) {
			return oaiClient{}, errors.New("azure openai deployment can only include alphanumeric characters, '_,-', and can't end with '_' or '-'")
		}
		azureClient, err = azureopenai.NewClient(*azureOpenAIEndpoint, *openAIAPIKey, *openAIDeploymentName)
		if err != nil {
			return oaiClients{}, fmt.Errorf("error create Azure client: %w", err)
		}
	}
	clients := oaiClients{
		azureClient:  azureClient,
		openAIClient: oaiClient,
	}

	return clients, nil
}

func completion(ctx context.Context, client oaiClients, prompts []string, deploymentName string, subCommand string) (string, error) {

}
