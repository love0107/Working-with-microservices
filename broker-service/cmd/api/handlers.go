package main

import (
	"encoding/json"
	"net/http"
)

type jsonResposne struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Date    any    `json:"data, omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResposne{
		Error:   false,
		Message: "Hit the broker",
	}
	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
