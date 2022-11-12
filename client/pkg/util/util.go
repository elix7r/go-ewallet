package util

import (
	"encoding/json"
)

func Map2Json(m map[string]interface{}) (string, error) {
	jsn, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return "", err
	}
	return string(jsn), nil
}

func Bytes2Json(bytes []byte) (string, error) {
	//var arr := make(s, 0)
	//
	//for _, e := range s {
	//	m, err := util2.Struct2Map(e)
	//	if err != nil {
	//		return "", err
	//	}
	//
	//	arr = append(arr, m)
	//}
	//
	//var res [][]byte
	//
	//for _, e = range arr {
	//	jsn, err := json.MarshalIndent(arr, "", "    ")
	//	if err != nil {
	//		return "", err
	//	}
	//
	//	res = append(res, jsn)
	//}

	prettyJson, err := json.MarshalIndent(bytes, "", "    ")
	if err != nil {
		return "", err
	}

	return string(prettyJson), nil

}
