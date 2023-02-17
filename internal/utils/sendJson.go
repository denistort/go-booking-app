package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJson(data interface{}, w http.ResponseWriter) {
	out, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}
