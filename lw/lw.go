package lw

import (
	"net/http"
	"polaroid/server"
	"polaroid/tool"
	"polaroid/types"
)

type Lw struct {
	Data *types.Data
}

func NewLw(s *types.Data) (a *Admin) {
	a = new(Admin)
	a.Data = s
	return
}

func (a *Lw) Fortune(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *Lw) Toilet(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *Lw) Rig(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *Lw) Figlet(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *Lw) Cow(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *Lw) WWW(s *server.Server) {
	s.NewR("/fortune", "fortune", []string{"GET"}, nil, 1)
	s.NewR("/toilet", "toilet", []string{"GET"}, nil, 1)
	s.NewR("/rig", "rig", []string{"GET"}, nil, 1)
	s.NewR("/figlet", "figlet", []string{"GET"}, nil, 1)
	s.NewR("/cow", "cow", []string{"GET"}, nil, 1)
}
