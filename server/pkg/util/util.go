package util

import (
	"encoding/json"
)

func Struct2Map(s interface{}) (map[string]interface{}, error) {
	var out map[string]interface{}

	data, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &out)
	return out, err
}

func Map2Struct(m map[string]interface{}, s interface{}) error {
	dbByte, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(dbByte, &s); err != nil {
		return err
	}

	return nil
}
