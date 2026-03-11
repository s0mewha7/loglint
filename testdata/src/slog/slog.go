package slog

import "log/slog"

func testLogs() {
	slog.Info("Starting server")

	slog.Info("запуск сервера")

	slog.Info("server started!!!")

	password := "123"
	slog.Info("password: " + password)

	slog.Info("server started")

}
