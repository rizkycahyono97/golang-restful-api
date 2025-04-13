package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequest(r *http.Request, result interface{}) {
	//Decode JSON Request Body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)
	PanicIfError(err)
}
