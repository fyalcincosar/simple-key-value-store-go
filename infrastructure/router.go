package infrastructure

import (
	"kvstore/interfaces"
	"kvstore/usecases"
	"net/http"
	"os"
)

func Dispatcher(logger usecases.Logger, handler interfaces.StoreHandler) {
	storeController := interfaces.NewStoreController(handler, logger)
	storeController.Save()
	mux := http.ServeMux{}
	mux.HandleFunc("/add", storeController.NewRecord)
	mux.HandleFunc("/get", storeController.GetValue)
	mux.HandleFunc("/set", storeController.SetValue)

	http.ListenAndServe(":"+os.Getenv("PORT"), &mux)
}
