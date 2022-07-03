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
	LastName string
	FirstName string
	Street string
	Country string
	Postal string
	Tel string
}


func GetName(a string) (e string, ee string) {
	fmt.Println(a)
	a = strings.TrimSpace(a);
	name := strings.Split(a, " ")
	fmt.Println(name, len(name))
	e = name[0]
	if len(ee) > 1 {
		ee = name[1]
	}
	return;
}

func MakeStr(e bytes.Buffer) (r []RigStruct) {
	a := e.String()
	aa := strings.TrimSpace(a)
	aaa := strings.Split(aa, "\n\n");
	for _, val := range aaa {
		var ee RigStruct
		v :=  strings.TrimSpace(val)
		vv := strings.Split(v, "\n")
		ee.FirstName, ee.LastName = GetName(vv[0])
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
		MakeStr(a);
		head := tool.NewHeader(r, w, "gopiko-rig", e)
		head.SetData(&Fortune{
			Fortune:a.String(),
		})
		head.Jointure("layout.html", "rig.html")
	})
}
