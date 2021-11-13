package infrastructure

import (
	"os"
	"testing"
)

func TestStoreHandler_Save(t *testing.T) {
	file, err := os.OpenFile("../temp/kvs_data_test.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(file.Name())
	testString := "SAVE TEST"
	err = os.WriteFile("../temp/kvs_data_test.json", []byte(testString), 066)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLoad(t *testing.T) {
	file, err := os.OpenFile("../temp/kvs_data_test.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(file.Name())

	_, err = os.ReadFile("../temp/kvs_data_test.json")
	if err != nil {
		t.Fatal(err)
	}
}
