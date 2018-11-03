package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type S struct {
	Profit    int64 `json:"profit"`
	Loss      int64 `json:"loss"`
	Affiliate int64 `json:"affiliate"`
}

var Length = 1000
var inc = 0

func Handler(w http.ResponseWriter, r *http.Request) {
	data := []S{}
	for i := 0; i < Length; i++ {
		data = append(data, S{
			Profit:    2000,
			Loss:      3000,
			Affiliate: 1000,
		})
	}
	inc++
	fmt.Println(inc)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
