package infrastructure

import (
	"encoding/json"
	"kvstore/domain"
	"log"
	"os"
	"time"
)

type StoreHandler struct {
	store *domain.Store
}

func NewStoreHandler(store *domain.Store) *StoreHandler {
	return &StoreHandler{
		store: store,
	}
}

func (sh *StoreHandler) GetRecord(key string) (value string, err error) {
	return sh.store.Get(key)
}

func (sh *StoreHandler) NewRecord(key, value string) error {
	return sh.store.Put(key, value)
}
func (sh *StoreHandler) SetValue(key, value string) error {
	return sh.store.Set(key, value)
}
func (sh *StoreHandler) Save() error {

	backupTime, err := time.ParseDuration(os.Getenv("BACKUP_TIME"))

	if err != nil {
		log.Printf("%s", err)
	}
	for range time.Tick(backupTime * 1) {
		file, err := os.OpenFile("./temp/kvs_data.json", os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Printf("%s", err)
			return err
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
			}
		}(file)

		records := domain.Records{}

		record := domain.Record{}

		for k, v := range sh.store.GetDb() {
			record.Key = k
			record.Value = v
			records = append(records, record)
		}
		data, err := json.Marshal(records)

		err = os.WriteFile("./temp/kvs_data.json", data, 066)
		if err != nil {
			return err
		}
	}
	return err
}
func (sh *StoreHandler) LoadStoreFromTemp() {

	file, e := os.ReadFile("./temp/kvs_data.json")

	records := domain.Records{}

	if e != nil {

		log.Printf("%s", e)

	} else {
		err := json.Unmarshal(file, &records)

		if err != nil {
			log.Printf("%s", err)
		}

		for _, record := range records {
			err := sh.NewRecord(record.Key, record.Value)
			if err != nil {
				return
			}
		}
	}
}
