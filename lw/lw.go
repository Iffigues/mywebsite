package lw

import (
	"net/http"
	"fmt"
	"polaroid/server"
	"polaroid/types"
	"polaroid/tool"
	"polaroid/linuxtool"
)

type Lw struct {
	Data *types.Data
}

func NewLw(s *types.Data) (a *Lw) {
	a = new(Lw)
	a.Data = s
	return
}

type Fortune struct {
	Fortune string
	Data []string
}

func (a *Lw) Fortune(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tt []linuxtool.Haha
		var rr  *linuxtool.Mimi = nil
		var err error
		var ttl linuxtool.Commande
		ll := ttl.GetFo()
		if r.Method == "GET" {
			ee := linuxtool.Haha{"-a", nil}
			tt = append(tt, ee)
			rr, err = e.Commande.Make("fortune", tt, nil)
			if err != nil {}
		}
		if r.Method == "POST" {
			aa := e.Commande.MakeFoArg(r)
			dd := e.Commande.MakeHaha("fortune", r)
			rr, err = e.Commande.Make("fortune", dd, aa)
			if err != nil {
				fmt.Println(err)
			}
		}
		a,b,c := e.Commande.Exec(rr)
		if c != nil {
		}
		fmt.Println(a, b , c)
		head := tool.NewHeader(r, w, "gopiko-fortune", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
			Data: ll,
		})
		head.Jointure("layout.html", "fortune.html")
	})
}

func (a *Lw) Toilet(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tt []linuxtool.Haha
		if r.Method == "GET" {
			ee := linuxtool.Haha{"-F",[]string{"gay"}}
			tt = append(tt, ee)
			oui := linuxtool.Haha{"-f",[]string{"mono12"}}
			tt = append(tt, oui)
			rr , err := e.Commande.Make("toilet", tt, []string{"welcome home"})
			if err != nil {
			}
			a,b,c := e.Commande.Exec(rr)
			if c != nil {
			}
			fmt.Println(a.String(), b.String() , c)
			head := tool.NewHeader(r, w, "gopiko-toilet", e)
			head.SetData(&Fortune{
				Fortune:a.String(),
			})
			head.Jointure("layout.html", "toilet.html")
		}
		if r.Method == "POST" {
		}
	})
}

func (a *Lw) Rig(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rr *linuxtool.Mimi
		var err error
		if r.Method == "GET" {
			rr , err = e.Commande.Make("rig", nil, nil)
			if err != nil {
			}
		}
		if r.Method == "POST" {
			dd := e.Commande.MakeHaha("rig", r)
			rr, err = e.Commande.Make("rig", dd, nil)
			if err != nil {
				fmt.Println(err)
			}
		}
		a, b, c := e.Commande.Exec(rr)
		if c != nil {
		}
		fmt.Println(a.String(), b.String() , c)
		head := tool.NewHeader(r, w, "gopiko-rig", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
		})
		head.Jointure("layout.html", "rig.html")
	})
}

func (a *Lw) Figlet(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rr  *linuxtool.Mimi = nil
		var err error

		if r.Method == "GET" {
			rr , err = e.Commande.Make("figlet", nil, []string{"welcome home"})
			if err != nil {
			}
		}
		if r.Method == "POST" {
		}
		a,b,c := e.Commande.Exec(rr)
		if c != nil {
		}
		fmt.Println(a.String(), b.String() , c)
		head := tool.NewHeader(r, w, "gopiko-rig", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
		})
		head.Jointure("layout.html", "figlet.html")
	})
}

func (a *Lw) Cow(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rr *linuxtool.Mimi = nil
		var err error
		var tt linuxtool.Commande
		ll := tt.GetAn()
		if r.Method == "GET" {
			rr , err = e.Commande.Make("cowsay", nil, []string{"welcome home"})
			if err != nil {
			}
		}
		if r.Method == "POST" {
			dd := e.Commande.MakeHaha("cowsay", r)
			rr, err = e.Commande.Make("cowsay", dd, []string{"zzz"})
			if err != nil {
				fmt.Println(err)
			}
		}
		a, _, c := e.Commande.Exec(rr)
		if c != nil {
			fmt.Println(c)
		}
		head := tool.NewHeader(r, w, "gopiko-cow", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
			Data:ll,
		})
		head.Jointure("layout.html", "cow.html")
	})
}

func (a *Lw) WWW(s *server.Server) {
	s.NewR("/fortune", "fortune", []string{"GET", "POST"}, a.Fortune(s.Data), 1)
	s.NewR("/toilet", "toilet", []string{"GET"}, a.Toilet(s.Data), 1)
	s.NewR("/rig", "rig", []string{"GET", "POST"}, a.Rig(s.Data), 1)
	s.NewR("/figlet", "figlet", []string{"GET"}, a.Figlet(s.Data), 1)
	s.NewR("/cowsay", "cowsay", []string{"GET","POST"}, a.Cow(s.Data), 1)
}
