package openrouter

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ClientResponse interface {
	GetResponse(prompt string) (string, error)
}

type Client struct {
	APIKey string
	APIUrl string
	Model  string
	Prompt string
}

func NewClient(APIKey, APIUrl, Model string, prompt string) *Client {
	return &Client{
		APIKey: APIKey,
		APIUrl: APIUrl,
		Model:  Model,
		Prompt: prompt,
	}
}
