package uberzap

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()

	// заглавная буква
	logger.Info("Starting server on port 8080")
	logger.Error("Failed to connect to database")
	logger.Warn("Retrying connection")
	logger.Debug("Processing request")

	// не английский
	logger.Info("запуск сервера")
	logger.Error("ошибка подключения")
	logger.Warn("повторная попытка")

	// спецсимволы и эмодзи
	logger.Info("server started!🚀")
	logger.Error("connection failed!!!")
	logger.Warn("something went wrong...")
	logger.Debug("request body: {id: 1}")

	// чувствительные данные
	password := "supersecret"
	token := "eyJhbGciOiJIUzI1NiJ9"
	apiKey := "sk-1234567890"

	logger.Info("user login: " + password)
	logger.Debug("auth header: " + token)
	logger.Info("api key: " + apiKey)

	// норм
	logger.Info("starting server on port 8080")
	logger.Error("failed to connect to database")
	logger.Warn("retrying connection")
	logger.Debug("processing request")
	logger.Info("user authenticated successfully")
	logger.Info("request completed")
}
