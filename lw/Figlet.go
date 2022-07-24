package lw

import (
	"net/http"
	"fmt"
	"polaroid/types"
	"polaroid/tool"
	"polaroid/linuxtool"
)

func (a *Lw) Figlet(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rr  *linuxtool.Mimi = nil
		var err error
		var tls linuxtool.Commande
		ll, lll := tls.GetFi()
		if r.Method == "GET" {
			rr , err = e.Commande.Make("figlet", nil, []string{"welcome home"})
			if err != nil {
			}
		}
		if r.Method == "POST" {
			dd := e.Commande.MakeHaha("figlet", r)
			rr, err = e.Commande.Make("figlet", dd, tls.GetMessage(r))
			if err != nil {
				fmt.Println(err)
			}
		}
		a, _, c := e.Commande.Exec(rr)
		if c != nil {
			fmt.Println(c)
		}
		head := tool.NewHeader(r, w, "gopiko-rig", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
			Data: ll,
			Datas: lll,
		})
		head.Jointure("layout.html", "figlet.html")
	})
}
