package service

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

// 使用 OpenAI API 进行文本审核
func ReviewText(input string) (string, error) {
	if OPENAI_API_KEY == "" {
		return "", errors.New("empty env.OPENAI_API_KEY")
	}
	if OPENAI_MODEL == "" {
		return "", errors.New("empty env.OPENAI_MODEL")
	}

	// 初始化配置
	config := openai.DefaultConfig(OPENAI_API_KEY)
	if OPENAI_BASE_URL != "" {
		config.BaseURL = OPENAI_BASE_URL
	}

	// 创建客户端
	client := openai.NewClientWithConfig(config)

	// 构建消息
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: SYSTEM_PROMPT,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		},
	}

	// 创建请求
	req := openai.ChatCompletionRequest{
		Model:    OPENAI_MODEL,
		Messages: messages,
	}

	// 发送请求
	cxt := context.Background()
	resp, err := client.CreateChatCompletion(cxt, req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
