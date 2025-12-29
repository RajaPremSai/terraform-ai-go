package cli

import (
	"context"

	openai "github.com/PullRequestInc/go-gpt3"
	azureopenai "github.com/RajaPremSai/terraform-ai-go/pkg/gpt3"
)

type oaiClients struct {
	azureClient  azureopenai.Client
	openAIClient openai.Client
}

const userRole="user"

func newOAIClients(oaiClients,error){

}

func Completion(ctx context.Context,client oaiClients,prompts []string,deploymentName string,subCommand string)(string,error){
	
}