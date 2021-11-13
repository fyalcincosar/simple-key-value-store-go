package main

import (
	"kvstore/domain"
	"kvstore/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()
	infrastructure.Load(logger)

	handler := infrastructure.NewStoreHandler(domain.NewStore())
	handler.LoadStoreFromTemp()
	infrastructure.Dispatcher(logger, handler)

}
