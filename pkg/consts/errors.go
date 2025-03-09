package consts

// ошибки конфигурации
const (
	TelegramTokenIsAbsent    = "отсутствует TELEGRAM_TOKEN"
	OpenRouterTokenIsAbsent  = "отсутствует токен OpenRouter"
	OpenRouterUrlIsAbsent    = "отсутствует API_URL"
	OpenRouterModelIsAbsent  = "отсутствует название Модели"
	OpenRouterPromptIsAbsent = "отсутсвует промпт для Модели"
)

// ошибки open_router
const (
	ResponseBodyError    = "ошибка при чтении ответа"
	ApiRouterError       = "ошибка API"
	JSONParsingError     = "ошибка при парсинге JSON"
	EmptyAnswerByAIError = "нет ответа от ИИ"
)
