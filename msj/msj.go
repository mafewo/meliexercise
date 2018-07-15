package msj

import (
	"encoding/json"
	"net/http"
)

//Data error data
type Data struct {
	w       http.ResponseWriter
	message string
	status  int
}

//Set Entrypoint to ERR
func Set(w http.ResponseWriter, message string, status int) *Data {
	return &Data{
		w:       w,
		message: message,
		status:  status,
	}
}

//ReturnJSON Return a JSON error message
func (d *Data) ReturnJSON() {

	jsonOutput, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{d.message})
	d.w.Header().Set("Content-Type", "application/json")
	d.w.WriteHeader(http.StatusForbidden)
	d.w.Write(jsonOutput)
}
