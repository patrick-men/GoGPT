package API

import(
	"fmt"
	"context"
	openai "github.com/sashabaranov/go-openai"
	viper "github.com/spf13/viper"
)

func GetKey() (apiKey string) {
	// Set config file
    viper.SetConfigFile("./Prompt/config.yaml")
    
	// Try to read config file, return any errors
	if err := viper.ReadInConfig(); err != nil {
        fmt.Println("Error reading config file:", err)
        return
    }

	// Read api_key from config.yaml file
    apiKey = viper.GetString("api_key")
    
	// Return error if key is empty
	if apiKey == "" {
        fmt.Println("API_KEY not found in config file")
        return
    }

	return apiKey
}

func APIrequest(prompt string) (answer string) {
	// Define client
	client := openai.NewClient(GetKey())

	// Prompt
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	answer = resp.Choices[0].Message.Content
	return answer
}
