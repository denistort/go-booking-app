package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Wewe struct {
	Name string
	Age  int
}

func create(name string, age int) Wewe {
	wewe := Wewe{
		name,
		age,
	}
	return wewe
}
func ReturnJson(w http.ResponseWriter, _ *http.Request) {
	slice := []Wewe{
		create("Denis", 40),
		create("Maksim", 22),
		create("Winston", 33),
	}

	b, err := json.Marshal(slice)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(b))
}
