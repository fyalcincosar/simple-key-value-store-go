package usecases

type StoreInteractor struct {
	StoreRepository StoreRepository
}

func (si *StoreInteractor) FindValueByKey(key string) (value string, err error) {
	return si.StoreRepository.FindValueByKey(key)
}

func (si *StoreInteractor) NewRecord(key, value string) (err error) {
	return si.StoreRepository.AddNewRecord(key, value)
}
func (si *StoreInteractor) Save() (err error) {
	return si.StoreRepository.Save()
}

func (si *StoreInteractor) SetRecord(key, value string) (err error) {
	return si.StoreRepository.SetRecord(key, value)
}
