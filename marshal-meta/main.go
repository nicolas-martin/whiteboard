package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	meta := OperationMetadata{
		Type: "a",
	}
	metaMap := map[string]interface{}{
		"type": "a2",
	}

	m2, _ := MarshalJSONMap(meta)

	bs := BigStruct{
		MetaData:    m2,
		MetaDataMap: metaMap,
	}

	b, _ := json.Marshal(bs)

	fmt.Println(string(b))

}

func UnmarshalJSONMap(m map[string]interface{}, i interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, i)
}

func MarshalJSONMap(i interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, nil
}

type OperationMetadata struct {
	Type string `json:"type"`
}

type BigStruct struct {
	MetaData    map[string]interface{} `json:"metadata"`
	MetaDataMap map[string]interface{} `json:"metadataMap"`
}
