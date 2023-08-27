package response

import "net/http"

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func Respond(w http.ResponseWriter, status int, data []byte) {
	SetJsonHeader(w)
	w.WriteHeader(status)
	w.Write(data)
}
