package uberzap

import "go.uber.org/zap"

func testZap() {
	logger, _ := zap.NewProduction()

	logger.Info("Server started")

	logger.Error("connection failed!!!")

	token := "abc123"

	logger.Info("token=" + token)

	logger.Info("connection established")

}
