package lw

import (
	"net/http"
	"fmt"
	"polaroid/types"
	"polaroid/tool"
	"polaroid/linuxtool"
)


func (a *Lw) Toilet(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tt []linuxtool.Haha
		var rr *linuxtool.Mimi = nil
		var err error
		var ttl linuxtool.Commande
		var i string = "1"

		if r.Method == "GET" {
			tt = append(tt, linuxtool.Haha{"-E", []string{"html"}})
			ee := linuxtool.Haha{"-F",[]string{"gay"}}
			tt = append(tt, ee)
			oui := linuxtool.Haha{"-f",[]string{"mono12"}}
			tt = append(tt, oui)
			rr , err = e.Commande.Make("toilet", tt, []string{"welcome home"})
			if err != nil {
			}
		}
		if r.Method == "POST" {
			tt = e.Commande.MakeHaha("toilet", r)
			fmt.Println(tt)
			rr, err = e.Commande.Make("toilet", tt, ttl.GetMessage(r))
			if err != nil {
				fmt.Println(err)
			}
		}
		a, _, c := e.Commande.Exec(rr)
		if c != nil {
		}
		rts, rtt := ttl.GetTo()
		ree, _ := ttl.GetFi()
		head := tool.NewHeader(r, w, "gopiko-toilet", e)
		i = ttl.GetTT(tt)
		println("i==",i)
		head.SetData(&Fortune{
			Fortune:a.String(),
			Data: rts,
			Datas: rtt,
			Datal: ree,
			I: i,
		})
		head.Jointure("layout.html", "toilet.html")
	})
}
