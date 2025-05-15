package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
)

func TestAi(c *gin.Context) {
	llm, err := openai.New(
		openai.WithModel("deepseek-reasoner"),
		openai.WithToken("sk-dd54e9f4bf6c43f9a103f39965b1e008"),
		openai.WithBaseURL("https://api.deepseek.com"),
	)
	if err != nil {
		log.Fatal("初始化失败:", err)
	}

	// 第二步：通过大模型llm调用API
	ctx := context.Background()
	println("早上好")
	prompt := "早上好"
	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		prompt,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 输出大模型生成的答复
	fmt.Println("这里是回复")
	fmt.Println(completion)
	//————————————————
	//
	//版权声明：本文为博主原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接和本声明。
	//
	//原文链接：https://blog.csdn.net/2201_75692232/article/details/146470874
}
