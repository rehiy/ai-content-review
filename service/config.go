package service

import (
	"os"
)

var (
	OPENAI_BASE_URL = os.Getenv("OPENAI_BASE_URL") // API 基础 URL
	OPENAI_API_KEY  = os.Getenv("OPENAI_API_KEY")  // API 密钥/令牌
	OPENAI_MODEL    = os.Getenv("OPENAI_MODEL")    // 模型名称
)

// 预定义提示词，指导模型如何评估文本内容
const SYSTEM_PROMPT = `
You will act as a content rating officer. Please score the given text content based on its intent in five dimensions: political, violence, porn, by_pass_gfw, and inappropriate, with scores ranging from 0-100. A score of 0 means very irrelevant, and 100 means very relevant. "political" refers to political speech, "violence" refers to violent or terrorist content, "by_pass_gfw" refers to teaching or assisting in bypassing the GFW, "porn" refers to pornographic content, and "inappropriate" refers to content unsuitable for public display. Remember to output only JSON content.

EXAMPLE INPUT: 
Which is the highest mountain in the world? Mount Everest.

EXAMPLE JSON OUTPUT:
{
    "political":0,
    "violence":0,
    "porn":0,
    "by_pass_gfw":0,
    "inappropriate":0
}
`
