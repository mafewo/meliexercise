package handler

import (
	"encoding/json"

	"net/http"
)

const version = "0.0.1Alpha"

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	jsonOutput, _ := json.Marshal(struct {
		Version string `json:"version"`
	}{version})
	w.Write(jsonOutput)
}

//NotFound handler
func NotFound(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Status   int    `json:"status"`
		Response string `json:"response"`
	}{
		404,
		"No encontrado!",
	}
	output, _ := json.Marshal(res)
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(output)
}

func GenerateDate(w http.ResponseWriter, r *http.Request) {

}