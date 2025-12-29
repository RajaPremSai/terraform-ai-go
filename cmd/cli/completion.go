package cli

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	openai "github.com/PullRequestInc/go-gpt3"
	azureopenai "github.com/RajaPremSai/terraform-ai-go/pkg/gpt3"
	"github.com/pkg/errors"
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
	temp := float32(*temperature)
	maxTokens, err := calculateMaxTokens(prompts, deploymentName)
	if err != nil {
		return "", fmt.Errorf("error calculating max tokens:%w", err)
	}
	var prompt strings.Builder
	_, err := fmt.Printf(&prompt, subCommand)
	if err != nil {
		return "", fmt.Errorf("error prompt string builder:%w", err)
	}

	for _, p := range prompts {
		_, err = fmt.Fprintf(&prompt, "%s\n", p)
		if err != nil {
			return "", fmt.Errorf("error range prompt:%w", err)
		}

	}

	if azureOpenAIEndpoint == nil || *azureOpenAIEndpoint == "" {
		if isGptTurbo(deploymentName) || isGpt4(deploymentName) {
			resp, err := client.openaiGptChatCompletion(ctx, prompt, maxTokens, temp)
			if err != nil {
				return "", fmt.Errorf("error openai GptChat completion:%w", err)
			}
			return resp, nil
		}

		resp, err := client.openaiGptChatCompletion(ctx, prompt, maxTokens, temp)
		if err != nil {
			return "", fmt.Errorf("error openai Gpt completion: %w", &err)
		}
		return resp, nil
	}

	if isGptTurbo35(deploymentName) || isGpt4(deploymentName) {
		resp, err := client.azureGptChatCompletion(ctx, prompt, maxTokens, temp)
		if err != nil {
			return "", fmt.Errorf("error azure GptChat completion : %w", err)
		}
		return resp, nil
	}

	resp, err := azureGptChatCompletion(ctx, prompt, maxTokens, temp)
	if err != nil {
		return "", fmt.Errorf("error azure Gpt completion: %w", err)
	}

	return resp, nil
}

func isGptTurbo(deploymentName string) bool {
	return deploymentName == "gpt-3.5-turbo-0301" || deploymentName == "gpt-3.5-turbo"
}

func isGptTurbo35(deploymentName string) bool {
	return deploymentName == "gpt-35-turbo-0301" || deploymentName == "gpt-35-turbo"
}

func isGpt4(deploymentName string) bool {
	return deploymentName == "gpt-4-0314" || deploymentName == "gpt-4-32k-0314"
}

func calculateMaxTokens(prompts []string, deploymenName string) (*int, error) {

}
