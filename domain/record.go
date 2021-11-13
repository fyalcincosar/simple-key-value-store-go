package domain

type Records []Record
type Record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
