package lw

import (
	"net/http"
	"fmt"
	"polaroid/types"
	"polaroid/tool"
	"polaroid/linuxtool"
)

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
			ee := tt.Think(r)
			fmt.Println(ee)
			rr, err = e.Commande.Make(ee, dd, tt.GetMessage(r))
			if err != nil {
				fmt.Println(err)
			}
		}
		a, _, c := e.Commande.Exec(rr)
		if c != nil {
		}
		head := tool.NewHeader(r, w, "gopiko-cow", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
			Data:ll,
		})
		head.Jointure("layout.html", "cow.html")
	})
}
