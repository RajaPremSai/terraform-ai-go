package cli

import (
	"context"
	"strings"

	"github.com/pkg/errors"
)

var errResp = errors.New("inavlid response")

func (c *oaiClients) openaiGptCompletion(ctx context.Context, prompt strings.Builder, maxTokens *int, temp float32) (string, error)

func (c *oaiClients) openaiGptChatCompletion(ctx context.Context, prompt strings.Builder, maxTokens *int, temp float32) (string, error)

func(c *oaiClients) azureGptCompletion(ctx context.Context, prompt strings.Builder, maxTokens *int, temp float32) (string, error)

func(c *oaiClients) azureChatGptCompletion(ctx context.Context, prompt strings.Builder, maxTokens *int, temp float32) (string, error)
