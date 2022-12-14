package addJsonToStruct

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func CustomUnmarshal(Struct any, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(file)
	if err = file.Close(); err != nil {
		return err
	}
	S := Struct
	if err := json.Unmarshal(byteValue, &S); err != nil {
		return err
	}
	return nil
}
