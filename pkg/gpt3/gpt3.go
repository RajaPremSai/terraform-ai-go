package gpt3

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultAPIVersion     = "2023-03-15-preview"
	defaultUserAgent      = "kubectl-openai"
	defaultTimeoutSeconds = 30
)

type Client interface {
	// ChatCompletion creates a completion with the Chat completion endpoint which
	// is what powers the ChatGPT experience.
	ChatCompletion(ctx context.Context, request ChatCompletionRequest) (*ChatCompletionResponse, error)

	// Completion creates a completion with the default engine. This is the main endpoint of the API
	// which auto-completes based on the given prompt.
	Completion(ctx context.Context, request CompletionRequest) (*CompletionResponse, error)

	// CompletionStream creates a completion with the default engine and streams the results through
	// multiple calls to onData.
	CompletionStream(ctx context.Context, request CompletionRequest, onData func(*CompletionResponse)) error

	// Given a prompt and an instruction, the model will return an edited version of the prompt.
	Edits(ctx context.Context, request EditsRequest) (*EditsResponse, error)

	// Search performs a semantic search over a list of documents with the default engine.
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Returns an embedding using the provided request.
	Embeddings(ctx context.Context, request EmbeddingsRequest) (*EmbeddingsResponse, error)
}

type client struct {
	endpoint       string
	apiKey         string
	deploymentName string
	apiVersion     string
	userAgent      string
	httpClient     *http.Client
}

func NewClient(endpoint string, apiKey string, deploymentName string, options ...ClientOption) (Client, error) {
	// Create a new HTTP client with a default timeout.
	httpClient := &http.Client{
		Timeout: defaultTimeoutSeconds * time.Second,
	}

	// Create a new client instance with the provided parameters.
	c := &client{
		endpoint:       endpoint,
		apiKey:         apiKey,
		deploymentName: deploymentName,
		apiVersion:     defaultAPIVersion,
		userAgent:      defaultUserAgent,
		httpClient:     httpClient,
	}

	// Apply any additional client options provided.
	for _, o := range options {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *client) Completion(ctx context.Context, request CompletionRequest) (*CompletionResponse, error) {
}

func (c *client) ChatCompletion(ctx context.Context, request ChatCompletionRequest) (*ChatCompletionResponse, error) {
}


var (
	dataPrefix   = []byte("data: ")
	doneSequence = []byte("[DONE]")
)

func (c *client) CompletionStream(ctx context.Context, request CompletionRequest, onData func(*CompletionResponse)) error {
}

func (c *client) Edits(ctx context.Context, request EditsRequest) (*EditsResponse, error) {
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
}

func (c *client) Embeddings(ctx context.Context, request EmbeddingsRequest) (*EmbeddingsResponse, error) {
}

func (c *client) performRequest(req *http.Request) (*http.Response, error) {
}

func checkForSuccess(resp *http.Response) error {
}

func getResponseObject(rsp *http.Response, v interface{}) error {
}

func jsonBodyReader(body interface{}) (io.Reader, error) {
}
 

func (c *client) newRequest(ctx context.Context, method, path string, payload interface{}) (*http.Request, error) {
}
