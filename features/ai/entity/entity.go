package entity

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type ChatRequest struct {
	Context  context.Context
	Client   *openai.Client
	Messages []openai.ChatCompletionMessage
	Model    string
}

type UseCaseInterface interface {
	RecommendRecyclable(itemName string) (string, error)
	GetCompletionFromMessages(request ChatRequest) (openai.ChatCompletionResponse, error)
}
