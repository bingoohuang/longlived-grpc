package handler

import (
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
	"regexp"
)

var (
	reGetFuncArg = regexp.MustCompile(`\( (.*) \) returns`)
	IndexHTML    *template.Template
	IndexFS      embed.FS
)

// Response - Standar ajax Response
type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data"`
}

func writeError(w http.ResponseWriter, err error) {
	e, _ := json.Marshal(Response{Error: err.Error()})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(e)
}

func response(w http.ResponseWriter, data interface{}) {
	e, _ := json.Marshal(Response{Data: data})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(e)
}
