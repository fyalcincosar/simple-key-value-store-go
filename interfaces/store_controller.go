package interfaces

import (
	"encoding/json"
	"kvstore/usecases"
	"net/http"
)

type StoreController struct {
	StoreInteractor usecases.StoreInteractor
	Logger          usecases.Logger
}

func NewStoreController(storeHandler StoreHandler, logger usecases.Logger) *StoreController {
	return &StoreController{
		StoreInteractor: usecases.StoreInteractor{
			StoreRepository: &StoreRepository{
				StoreHandler: storeHandler,
			},
		},
		Logger: logger,
	}
}

func (sc *StoreController) GetValue(w http.ResponseWriter, r *http.Request) {

	sc.Logger.LogHttpServer("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	key := r.URL.Query().Get("key")

	value, err := sc.StoreInteractor.FindValueByKey(key)

	if err != nil {

		sc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(204)

		jsonErr := json.NewEncoder(w).Encode(err.Error())

		if jsonErr != nil {

			sc.Logger.LogError("%s", jsonErr)

		}

	} else {

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(200)

		record := struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{key, value}

		jsonErr := json.NewEncoder(w).Encode(record)

		if jsonErr != nil {

			sc.Logger.LogError("%s", jsonErr)

		}
	}
}

func (sc *StoreController) NewRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		sc.Logger.LogHttpServer("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		key := r.URL.Query().Get("key")

		value := r.URL.Query().Get("value")

		err := sc.StoreInteractor.NewRecord(key, value)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {

			sc.Logger.LogError("%s", err)

			w.WriteHeader(204)

			jsonErr := json.NewEncoder(w).Encode(err.Error())

			if jsonErr != nil {

				sc.Logger.LogError("%s", jsonErr)

			}
		}

	} else {

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(404)

	}
}
func (sc *StoreController) SetValue(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		sc.Logger.LogHttpServer("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		key := r.URL.Query().Get("key")

		value := r.URL.Query().Get("value")

		err := sc.StoreInteractor.SetRecord(key, value)

		if err != nil {

			sc.Logger.LogError("%s", err)

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(204)

			jsonErr := json.NewEncoder(w).Encode(err)

			if jsonErr != nil {

				sc.Logger.LogError("%s", jsonErr)

			}

		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(200)

		record := struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{key, value}

		jsonErr := json.NewEncoder(w).Encode(record)

		if jsonErr != nil {

			sc.Logger.LogError("%s", jsonErr)

		}

	} else {

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(404)

	}
}

func (sc *StoreController) Save() {

	go sc.StoreInteractor.Save()

}
