package handler

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Init - routes initialization
func Init(router *mux.Router) {
	h := InitHandler()

	router.HandleFunc("/", h.index)

	ajaxRoute := router.PathPrefix("/server/{host}").Subrouter()
	ajaxRoute.HandleFunc("/services", cors(h.getLists)).Methods(http.MethodGet, http.MethodOptions)
	ajaxRoute.HandleFunc("/services", cors2(h.getListsWithProto)).Methods(http.MethodPost)
	ajaxRoute.HandleFunc("/service/{serv_name}/functions", cors(h.getLists)).Methods(http.MethodGet, http.MethodOptions)
	ajaxRoute.HandleFunc("/function/{func_name}/describe", cors2(h.describeFunction)).Methods(http.MethodGet, http.MethodOptions)
	ajaxRoute.HandleFunc("/function/{func_name}/invoke", cors(h.invokeFunction)).Methods(http.MethodPost, http.MethodOptions)

	// get list of active connection
	router.HandleFunc("/active/get", cors(h.getActiveConns)).Methods(http.MethodGet, http.MethodOptions)
	// close active connection
	router.HandleFunc("/active/close/{host}", cors2(h.closeActiveConns)).Methods(http.MethodDelete, http.MethodOptions)

	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", sub("index/css")))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", sub("index/js")))
	router.PathPrefix("/font/").Handler(http.StripPrefix("/font/", sub("index/font")))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", sub("index/img")))
}

func sub(dir string) http.Handler {
	dir = strings.TrimSuffix(dir, "/")
	dir = strings.TrimPrefix(dir, "/")
	s, err := fs.Sub(IndexFS, dir)
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(s))
}

func cors2(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "use_tls")
			return
		}

		if err := h(w, r); err != nil {
			writeError(w, err)
		}
	}
}

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "use_tls")
			return
		}

		h.ServeHTTP(w, r)
	}
}
