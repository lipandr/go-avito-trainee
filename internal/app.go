package internal

import (
	"net/http"
	"time"
)

//type Application interface {
//	FormHandler(w http.ResponseWriter, r *http.Request)
//	ListHandler(w http.ResponseWriter, r *http.Request)
//}

type server struct {
	http.Server
	svc *service
}

func NewServer(cfg Config, svc *service) *server {
	LoadHTMLTemplates()
	s := &server{
		Server: http.Server{
			Addr:         cfg.ServerAddress,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		svc: svc,
	}
	http.HandleFunc("/", s.FormHandler)
	http.HandleFunc("/list", s.ListHandler)

	return s
}
