package interfaces

type StoreRepository struct {
	StoreHandler StoreHandler
}

func (sr *StoreRepository) FindValueByKey(key string) (string, error) {
	return sr.StoreHandler.GetRecord(key)
}
func (sr *StoreRepository) AddNewRecord(key, value string) error {
	return sr.StoreHandler.NewRecord(key, value)
}
func (sr *StoreRepository) Save() error {
	sr.StoreHandler.Save()
	return nil
}
func (sr *StoreRepository) SetRecord(key, value string) error {
	return sr.StoreHandler.SetValue(key, value)
}
