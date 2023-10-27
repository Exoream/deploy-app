package usecase

import (
	"context"
	"errors"
	"recycle/features/ai/entity"

	openai "github.com/sashabaranov/go-openai"
)

type RubbishUseCase struct {
	userInput entity.UseCaseInterface
	openAIKey string
}

func NewAIUsecase(userInput entity.UseCaseInterface, openAIKey string) entity.UseCaseInterface {
	return &RubbishUseCase{
		userInput: userInput,
		openAIKey: openAIKey,
	}
}

// RecommendRecyclable implements entity.UseCaseInterface.
func (uc *RubbishUseCase) RecommendRecyclable(itemName string) (string, error) {
	chatSystem := "Selamat datang! Saya adalah sistem yang dapat memberikan informasi tentang jenis sampah yang bisa didaur ulang, tujuan penggunaan sampah daur ulangnya, dan fun fact tentang sampah tersebut."
	if itemName == "" {
		return "", errors.New("type field is required")
	}

	ctx := context.Background()
	client := openai.NewClient(uc.openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: chatSystem,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: itemName,
		},
	}

	request := entity.ChatRequest{
		Context:  ctx,
		Client:   client,
		Model:    model,
		Messages: messages,
	}

	resp, err := uc.GetCompletionFromMessages(request)
	if err != nil {
		return "", err
	}

	answer := resp.Choices[0].Message.Content
	return answer, nil
}

// GetCompletionFromMessages implements entity.UseCaseInterface.
func (uc *RubbishUseCase) GetCompletionFromMessages(request entity.ChatRequest) (openai.ChatCompletionResponse, error) {
	if request.Model == "" {
		request.Model = openai.GPT3Dot5Turbo
	}

	resp, err := request.Client.CreateChatCompletion(
		request.Context,
		openai.ChatCompletionRequest{
			Model:    request.Model,
			Messages: request.Messages,
		},
	)
	return resp, err
}
