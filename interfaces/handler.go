package interfaces

type StoreHandler interface {
	GetRecord(key string) (string, error)
	NewRecord(key, value string) error
	Save() error
	SetValue(key, value string) error
}
