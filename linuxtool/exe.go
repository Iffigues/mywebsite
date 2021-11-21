package linuxtool

import (
	"net/http"
	"errors"
	"bytes"
	"fmt"
	"log"
	"strings"
	"os/exec"
)

type Mimi struct {
	Com string
	Opt []string
}

type Haha struct {
	B string
	C []string
}

type How struct {
	A string
	B int
}

type Bash struct {
	Bash string
	Arg  bool
	Com  []How
	Args []string
	Obl []string
}

type Commande struct {
	Ani []string
	Fo []string
	To []string
	Toto []string
	Fi []string
	Fifi []string
	Co map[string]Bash
}

func (r *Commande) GetAn() (a []string){
	cmd := exec.Command("/usr/games/cowsay","-l")
	st, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	dd := strings.Trim(string(st), "Cow files in /usr/share/cowsay/cows:")
	dd = strings.TrimSpace(dd)
	return strings.Split(dd," ")
}

func (r *Commande) GetTo() (a []string, b []string){
	cmd := exec.Command("toilet","-F","list")
	st, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	ff := strings.Split(string(st),"\n")
	for _, val := range ff[1:len(ff) - 1] {
		r := strings.Split(val, "\"")
		a = append(a, r[1])
	}
	cmd = exec.Command("toilet", "-E", "list")
	st, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	ff = strings.Split(string(st), "\n")
	for _, val :=  range ff[1:len(ff) - 1] {
		r := strings.Split(val, "\"")
		b = append(b, r[1])
	}
	return
}


func (r *Commande) GetFi() (a []string, b []string){
	cmd := exec.Command("figlist")
	st, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	ll := strings.Split(string(st), "Figlet control files in this directory:")
	a = strings.Split(ll[0], "\n")[3:]
	b = strings.Split(ll[1], "\n")
	return
}

func (r *Commande) GetFo() (a []string){
	cmd := exec.Command("/usr/games/fortune","-f")
	st, err := cmd.CombinedOutput()
	if err != nil {
	}
	dd := strings.Trim(string(st), "100.00% /usr/share/games/fortunes")
	d := strings.Split(dd,"\n")
	d = d[1:len(d)-1]
	for k,v := range d {
		d[k] = strings.TrimSpace(v)
		e := strings.Split(d[k], " ")
		d[k] = e[1]
	}
	return d
}


func NewCommand() (r *Commande){
	r = &Commande{}
	r.Ani = r.GetAn()
	r.Fo = r.GetFo()
	r.To, r.Toto = r.GetTo()
	r.Fi, r.Fifi = r.GetFi()
	r.Co = make(map[string]Bash)
	r.Co["fortune"] = MakeFortune()
	r.Co["figlet"] = MakeFiglet()
	r.Co["toilet"] = MakeToilet()
	r.Co["cowsay"] = MakeCowsay()
	r.Co["cowthink"] = MakeCowthink()
	r.Co["rig"] = MakeRig()
	return
}


func (r *Commande) MakeFoArg(re *http.Request) (c []string){
	c = nil
	re.ParseForm()
	for k,v := range re.Form {
		fmt.Println(v)
		if strings.HasPrefix(k, "type-") {
			if vv, ok := re.Form["percent-" + k[5:]]; ok {
				fmt.Println("vv",vv)
			}
		}
	}
	return
}

func (r *Commande) MakeHaha(a string, re *http.Request) (b []Haha){
	re.ParseForm()
	tt := r.Co[a].Com
	for key, val := range re.Form{
			if len(val[0]) > 0 {
			var gg Haha
			fmt.Println("ez",key, val, len(val))
			for _, vals := range tt {
				if key == vals.A {
					gg.B = key
					if vals.B > 0 {
						gg.C  = append(gg.C, val[0])
					}
					b = append(b, gg)
				}
			}
		}
	}
	fmt.Println(b)
	return
}

func (r *Commande) Make(a string, b []Haha, c []string) (m *Mimi, err error){
	if val, ok := r.Co[a]; ok {
		m = &Mimi{}
		if val.Arg  && c == nil {
			return
		}
		for _, val := range val.Obl {
			m.Opt =  append(m.Opt, val)
		}
		m.Com = val.Bash
		for _, vali := range b {
			for _, k := range val.Com {
				if k.A == vali.B && k.B == len(vali.C) {
					m.Opt = append(m.Opt, vali.B)
					for _, rt := range vali.C {
						m.Opt = append(m.Opt, rt)
					}
				}
			}
		}
		for _, val := range c {
			m.Opt = append(m.Opt, val)
		}
		return m, nil

	}
	return nil,    errors.New("no command found")
}

func (r *Commande) Exec(m *Mimi) (out, er bytes.Buffer, err error) {
	cmd := exec.Command(m.Com, m.Opt...)
	cmd.Stdout = &out
	cmd.Stderr  = &er
	err = cmd.Run()
	if err != nil {
		return out, er, err
	}
	return
}
