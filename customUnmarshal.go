package addJsonToStruct

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func CustomUnmarshal(Struct any, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(file)
	if err = file.Close(); err != nil {
		panic(err)
	}
	S := Struct
	if err := json.Unmarshal(byteValue, &S); err != nil {
		panic(err)
	}
}
