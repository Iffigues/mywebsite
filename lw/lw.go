package lw

import (
	"net/http"
	"polaroid/server"
	"polaroid/types"
	"polaroid/tool"
)

func (a *Lw) Ascii(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := tool.NewHeader(r, w, "gopiko-ascii", e)
		head.Jointure("layout.html", "ascii.html")
	})
}

func (a *Lw) WWW(s *server.Server) {
	s.NewR("/fortune", "fortune", []string{"GET", "POST"}, a.Fortune(s.Data), 1)
	s.NewR("/toilet", "toilet", []string{"GET", "POST"}, a.Toilet(s.Data), 1)
	s.NewR("/rig", "rig", []string{"GET", "POST"}, a.Rig(s.Data), 1)
	s.NewR("/figlet", "figlet", []string{"GET", "POST"}, a.Figlet(s.Data), 1)
	s.NewR("/cowsay", "cowsay", []string{"GET","POST"}, a.Cow(s.Data), 1)
	s.NewR("/ascii", "ascii", []string{"GET"}, a.Ascii(s.Data), 1)
}
