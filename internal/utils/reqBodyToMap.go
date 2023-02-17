package utils

import (
	"errors"
	"net/http"
)

func ReqBodyToMap(fields []string, req *http.Request) (map[string]string, error) {
	res := make(map[string]string)
	for i, value := range fields {
		if req.Form.Has(fields[i]) == false {
			return nil, errors.New("wewe")
		}
		res[value] = req.Form.Get(fields[i])
	}
	return res, nil
}
