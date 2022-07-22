package lw

import (
	"bytes"
	"net/http"
	"fmt"
	"strings"
	"polaroid/types"
	"polaroid/tool"
	"polaroid/linuxtool"
)

type RigStruct struct {
	Name string
	Addr	string
	Postal string
	Tel string
}

func MakeStr(e bytes.Buffer) (r []RigStruct) {
	a := e.String()
	aa := strings.TrimSpace(a)
	aaa := strings.Split(aa, "\n\n");
	for _, val := range aaa {
		var ee RigStruct
		v :=  strings.TrimSpace(val)
		vv := strings.Split(v, "\n")
		ee.Name = vv[0]
		ee.Addr = vv[1]
		ee.Postal = vv[2]
		ee.Tel = vv[3]
		r = append(r, ee)
	}
	return
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
		a, _, c := e.Commande.Exec(rr)
		if c != nil {
		}
		eee := MakeStr(a);
		head := tool.NewHeader(r, w, "gopiko-rig", e)
		head.SetData(eee)
		head.Jointure("layout.html", "rig.html")
	})
}
