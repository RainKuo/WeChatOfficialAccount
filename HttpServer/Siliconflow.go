package HttpServer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request发起结构
type ReqContent struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Stream           bool     `json:"stream"`
	MaxTokens        int      `json:"max_tokens"`
	Stop             []string `json:"stop"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"top_p"`
	TopK             int      `json:"top_k"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	N                int      `json:"n"`
	ResponseFormat   struct {
		Type string `json:"type"`
	} `json:"response_format"`
	Tools []struct {
		Type     string `json:"type"`
		Function struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Parameters  struct {
			} `json:"parameters"`
			Strict bool `json:"strict"`
		} `json:"function"`
	} `json:"tools"`
}

// Response返回接口
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Response struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

func CreateReqContent(msgContent string) string {
	content := &ReqContent{
		Model: "deepseek-ai/DeepSeek-V3",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: msgContent,
			},
		},
		Stream:           false,
		MaxTokens:        1024,
		Stop:             []string{"null"},
		Temperature:      0.7,
		TopP:             0.7,
		TopK:             50,
		FrequencyPenalty: 0.5,
		N:                1,
		ResponseFormat: struct {
			Type string `json:"type"`
		}{Type: "text"},
		Tools: []struct {
			Type     string `json:"type"`
			Function struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Parameters  struct {
				} `json:"parameters"`
				Strict bool `json:"strict"`
			} `json:"function"`
		}{
			{
				Type: "function",
				Function: struct {
					Description string   `json:"description"`
					Name        string   `json:"name"`
					Parameters  struct{} `json:"parameters"`
					Strict      bool     `json:"strict"`
				}{
					Description: "Sample function",
					Name:        "exampleFunction",
					Strict:      false,
				},
			},
		},
	}
	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonData)
}

func DoSiliconflowRequest() {
	url := "https://api.siliconflow.cn/v1/chat/completions"
	content := CreateReqContent("我创建了一个微信公众号,用户的对话会直接发送给deepseek的api,我应该如何设计服务器逻辑")
	// payload := strings.NewReader("{\n  \"model\": \"deepseek-ai/DeepSeek-V3\",\n  \"messages\": [\n    {\n      \"role\": \"user\",\n      \"content\": \"中国大模型行业2025年将会迎来哪些机遇和挑战？\"\n    }\n  ],\n  \"stream\": false,\n  \"max_tokens\": 512,\n  \"stop\": [\n    \"null\"\n  ],\n  \"temperature\": 0.7,\n  \"top_p\": 0.7,\n  \"top_k\": 50,\n  \"frequency_penalty\": 0.5,\n  \"n\": 1,\n  \"response_format\": {\n    \"type\": \"text\"\n  },\n  \"tools\": [\n    {\n      \"type\": \"function\",\n      \"function\": {\n        \"description\": \"<string>\",\n        \"name\": \"<string>\",\n        \"parameters\": {},\n        \"strict\": false\n      }\n    }\n  ]\n}")
	// fmt.Println(content)
	payload := strings.NewReader(content)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "Bearer sk-utbpkbsmzlqhmrobnuxgxdirtlgqnoqxqevryomxiswlhadh")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	if len(response.Choices) > 0 {
		fmt.Println(response.Choices[0].Message.Content)
	}
}
