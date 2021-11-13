package usecases

type StoreRepository interface {
	FindValueByKey(key string) (string, error)
	AddNewRecord(key, value string) error
	Save() error
	SetRecord(key, value string) error
}
