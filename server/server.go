package server

import (
	"log"
	"net/http"

	"github.com/Drambluker/godirect/config"
	"github.com/Drambluker/godirect/utils"
	"github.com/gorilla/mux"
)

const key = "key"

type Server struct {
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{config: config}
}

func (s *Server) Run() {
	conf := s.config
	scheme := conf.Scheme

	context := conf.ContextPath
	router := mux.NewRouter()
	router.
		HandleFunc("/"+context+"/{"+key+"}", makeRedirectHandler(context, conf.Rules)).
		Host(conf.Host).
		Methods("GET").
		Schemes(scheme)
	http.Handle("/", router)

	log.Printf("Server is listening...")

	if scheme == "https" {
		tls := conf.TLS
		log.Fatal(http.ListenAndServeTLS(utils.FormatAllNetAddr(conf.SecurePort), tls.CertFile, tls.KeyFile, nil))
	} else {
		log.Fatal(http.ListenAndServe(utils.FormatAllNetAddr(conf.Port), nil))
	}
}

func makeRedirectHandler(context string, rules map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		from := vars[key]
		to := rules[from]
		log.Printf("Redirecting from /%s/%s to %s", context, from, to)
		http.Redirect(w, r, to, http.StatusFound)
	}
}
