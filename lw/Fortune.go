package lw


import (
	"net/http"
	"fmt"
	"polaroid/types"
	"polaroid/tool"
	"polaroid/linuxtool"
)

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
		a, _, c := e.Commande.Exec(rr)
		if c != nil {
		}
		head := tool.NewHeader(r, w, "gopiko-fortune", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
			Data: ll,
		})
		head.Jointure("layout.html", "fortune.html")
	})
}
