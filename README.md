# Terraform AI Assistant

A powerful command-line tool written in Go that leverages OpenAI's GPT models to generate and apply Terraform configurations. This assistant helps you create Terraform HCL templates through natural language prompts, making infrastructure-as-code more accessible and efficient.

## Features

- 🤖 **AI-Powered Code Generation**: Generate Terraform configurations using OpenAI GPT models (GPT-3.5, GPT-4) or Azure OpenAI
- 🔄 **Interactive Workflow**: Review and approve generated templates before applying
- ✅ **Template Validation**: Automatically validates generated Terraform HCL before execution
- 🚀 **Seamless Integration**: Works with your existing Terraform installation
- 🔧 **Dual Commands**: Separate commands for provider initialization and resource creation
- 🌐 **Multi-Cloud Support**: Supports both OpenAI and Azure OpenAI endpoints

## Prerequisites

- Go 1.23.2 or higher
- Terraform installed and available in your PATH
- OpenAI API key or Azure OpenAI endpoint and API key

## Installation

### Build from Source

```bash
git clone https://github.com/RajaPremSai/terraform-ai-go.git
cd terraform-ai-go
go build -o terraform-assistant
```

### Using Go Install

```bash
go install github.com/RajaPremSai/terraform-ai-go@latest
```

## Configuration

### Environment Variables

The tool can be configured using environment variables or command-line flags:

| Variable | Flag | Description | Required |
|----------|------|-------------|----------|
| `OPENAI_API_KEY` | `--openai-api-key` | OpenAI API key | Yes* |
| `AZURE_OPENAI_ENDPOINT` | `--azure-openai-endpoint` | Azure OpenAI endpoint URL | No |
| `OPENAI_DEPLOYMENT_NAME` | `--openai-deplyment-name` | Model deployment name (default: `text-davinci-003`) | No |
| `WORKING_DIR` | `--working-dir` | Terraform working directory | No |
| `EXEC_DIR` | `--exe-dir` | Path to Terraform executable | No |
| `TEMPERATURE` | `--temperature` | Model temperature (default: `0.0`) | No |
| `MAX_TOKENS` | `--max-tokens` | Maximum tokens for completion | No |
| `REQUIRED_CONFIRMATION` | `--required-confirmation` | Require confirmation before applying (default: `true`) | No |

*Required if not using Azure OpenAI

### Supported Models

The tool supports the following models with optimized token limits:

- `code-davinici-002`: 4097 tokens
- `text-daavinci-003`: 4097 tokens
- `gpt-3.5-turbo-0301`: 4096 tokens
- `gpt-35-turbo-0301`: 4096 tokens (Azure)
- `gpt-4-0314`: 8192 tokens
- `gpt-4-32k-0314`: 8192 tokens

## Usage

### Basic Usage

```bash
# Set your OpenAI API key
export OPENAI_API_KEY="your-api-key-here"

# Generate and apply Terraform configuration
terraform-assistant "create an S3 bucket named my-bucket"
```

### Initialize Provider Configuration

The `init` command generates provider configuration templates:

```bash
terraform-assistant init "configure AWS provider with region us-east-1"
```

This will:
1. Generate a provider configuration template
2. Save it as `provider.tf`
3. Run `terraform init`

### Generate Resource Configuration

The default command generates resource configurations:

```bash
terraform-assistant "create an EC2 instance with t2.micro instance type"
```

This will:
1. Generate Terraform HCL for the requested resource
2. Prompt you to review and approve
3. Save the configuration to a `.tf` file
4. Run `terraform apply`

### Interactive Workflow

When you run a command, the tool will:

1. **Generate Template**: Use AI to create Terraform HCL based on your prompt
2. **Display Preview**: Show you the generated configuration
3. **User Confirmation**: Prompt you with options:
   - `Apply`: Save and apply the configuration
   - `Don't Apply`: Exit without applying
   - `Reprompt`: Regenerate with modifications
4. **Validation**: Validate the Terraform syntax
5. **Execution**: Apply the configuration if approved

### Using Azure OpenAI

```bash
export AZURE_OPENAI_ENDPOINT="https://your-resource.openai.azure.com"
export OPENAI_API_KEY="your-azure-api-key"
export OPENAI_DEPLOYMENT_NAME="gpt-35-turbo-0301"

terraform-assistant "create a resource group in Azure"
```

### Advanced Options

```bash
# Use a specific model
terraform-assistant --openai-deplyment-name gpt-4-0314 "create a VPC"

# Adjust temperature for more creative outputs
terraform-assistant --temperature 0.7 "create a complex infrastructure setup"

# Skip confirmation prompt
terraform-assistant --required-confirmation=false "create a simple resource"

# Specify custom working directory
terraform-assistant --working-dir /path/to/terraform/project "create resources"
```

## Architecture

### Project Structure

```
terraform-ai-go/
├── cmd/
│   └── cli/              # CLI command implementations
│       ├── completion.go # GPT completion logic
│       ├── init.go       # Init command handler
│       ├── openai.go     # OpenAI client implementations
│       ├── root.go       # Root command setup
│       ├── run.go        # Main run command handler
│       └── util.go       # Utility functions
├── pkg/
│   ├── gpt3/             # Azure OpenAI client implementation
│   ├── terraform/        # Terraform operations
│   │   ├── impl.go       # Terraform operation implementations
│   │   ├── ops.go        # Terraform operations interface
│   │   ├── terraform.go  # Terraform client wrapper
│   │   └── validator.go  # HCL validation
│   └── utils/            # Utility functions
│       ├── file.go       # File operations
│       ├── terraform.go  # Terraform utilities
│       └── utils.go      # General utilities
└── main.go               # Application entry point
```

### Core Functions

#### Main Function
- Initializes the working directory
- Sets up Terraform execution path
- Calls `InitAndExecute` to start the application

#### InitAndExecute
- First function called from main
- Validates OpenAI API key presence
- Parses command-line flags and environment variables
- Executes the root command

#### RootCmd
- Sets up the Cobra CLI command structure
- Configures flags and subcommands
- Initializes Terraform operations interface

#### Run Function
The heart of the business logic:
- Creates OpenAI clients (OpenAI or Azure OpenAI)
- Generates Terraform HCL using GPT completion
- Generates appropriate filename for the template
- Prompts user for approval
- Validates the generated template
- Saves the file and applies Terraform configuration

#### newOAIClients
- Creates and returns OpenAI client instances
- Supports both OpenAI and Azure OpenAI endpoints
- Validates deployment name format for Azure

#### completion (gptCompletion)
- Generates completions for given prompts using OpenAI GPT models
- Automatically selects between GPT completion and Chat GPT APIs
- Handles token calculation and model-specific limits
- Supports both OpenAI and Azure OpenAI services

#### openaiGptCompletion
- Sends completion requests to OpenAI GPT-3 API
- Returns generated text based on provided prompt
- Handles single-prompt completions

#### openaiChatGptCompletion
- Similar to `openaiGptCompletion` but uses Chat API
- Supports multiple messages in conversation format
- Works with GPT-3.5-turbo and GPT-4 models

#### userActionPrompt
- Interactive prompt to confirm generated Terraform manifest
- Options: Apply, Don't Apply, or Reprompt
- Can be disabled via `--required-confirmation` flag

#### applyManifest
- Applies the generated Terraform configuration
- Executes `terraform apply` command
- Shows spinner during execution

#### Terraform Operations
- **Init()**: Runs `terraform init` with spinner feedback
- **Apply()**: Runs `terraform apply` with spinner feedback
- **CheckTemplate()**: Validates Terraform HCL syntax

#### Utility Functions
- **GetName()**: Generates or validates Terraform filename
- **StoreFile()**: Saves generated templates to disk
- **TerraformPath()**: Locates Terraform executable in PATH
- **CurrentDir()**: Gets current working directory
- **calculateMaxTokens()**: Calculates optimal token limits based on prompt length

## Examples

### Example 1: Create AWS S3 Bucket

```bash
terraform-assistant "create an S3 bucket named my-app-bucket with versioning enabled"
```

### Example 2: Create Azure Resource Group

```bash
export AZURE_OPENAI_ENDPOINT="https://your-resource.openai.azure.com"
terraform-assistant "create an Azure resource group named rg-production in East US region"
```

### Example 3: Initialize AWS Provider

```bash
terraform-assistant init "configure AWS provider for us-west-2 region with default tags"
```

### Example 4: Complex Infrastructure

```bash
terraform-assistant "create a VPC with public and private subnets, internet gateway, NAT gateway, and security groups for web and database tiers"
```

## Error Handling

The tool includes comprehensive error handling:

- **API Key Validation**: Exits with clear error if API key is missing
- **Template Validation**: Validates HCL syntax before applying
- **Terraform Errors**: Propagates Terraform execution errors with context
- **Model Selection**: Validates deployment names and model compatibility
- **Token Limits**: Automatically calculates and enforces token limits

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI
- Uses [terraform-exec](https://github.com/hashicorp/terraform-exec) for Terraform operations
- Powered by OpenAI GPT models

## Support

For issues, questions, or contributions, please open an issue on the GitHub repository.

---

**Note**: Always review generated Terraform configurations before applying them to production environments. The AI-generated code should be treated as a starting point and may require manual adjustments.
