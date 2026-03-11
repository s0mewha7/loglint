package slog

import "log/slog"

func main() {
	// заглавная буква
	slog.Info("Starting server on port 8080")
	slog.Error("Failed to connect to database")
	slog.Warn("Retrying connection")
	slog.Debug("Processing request")

	// не английский
	slog.Info("запуск сервера")
	slog.Error("ошибка подключения к базе данных")
	slog.Warn("повторная попытка подключения")

	// спецсимволы и эмодзи
	slog.Info("server started!🚀")
	slog.Error("connection failed!!!")
	slog.Warn("something went wrong...")
	slog.Debug("request body: {id: 1}")

	// чувствительные данные
	password := "supersecret"
	token := "eyJhbGciOiJIUzI1NiJ9"
	apiKey := "sk-1234567890"

	slog.Info("user login: " + password)
	slog.Debug("auth: " + token)
	slog.Info("key: " + apiKey)

	// норм
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
	slog.Warn("retrying connection")
	slog.Debug("processing request")
	slog.Info("user authenticated successfully")
	slog.Info("request completed")
}
