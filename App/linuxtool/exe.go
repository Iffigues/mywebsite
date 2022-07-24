package linuxtool

import (
	"net/http"
	"errors"
	"bytes"
	"os"
	"fmt"
	"log"
	"strings"
	"os/exec"
	"time"
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

func (r Commande) GetTT(a []Haha) (i string) {
	i = "2"
	for _, v := range a {
		if v.B == "-E" {
			for _, k := range v.C {
				if k == "html" || k == "html3" {
					i = "1"
				}
			}
		}
	}
	return i
}

func (r *Commande) GetAn() (a []string){
	cmd := exec.Command("/usr/games/cowsay","-l")
	st, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	dd := strings.Trim(string(st), "Cow files in /usr/share/cowsay/cows:")
	dd = strings.TrimSpace(dd)
	ff :=  strings.Split(dd, " ")
	for _, gg := range ff {
		ffg := strings.Split(gg, "\n")
		for _, ez := range ffg {
			a = append(a, ez)
		}
	}
	return 
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
	fmt.Println("haha-",a, b)
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
	a[len(a) - 1] = "mono12"
	b = strings.Split(ll[1], "\n")[1:]
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
			if strings.HasPrefix(k, "type-") {
				if len(v) > 0 && v[0] != "" {
				if vv, ok := re.Form["percent-" + k[5:]]; ok {
					if len(vv) > 0 && vv[0] != ""  {
						c = append(c, vv[0] + "%")
					}
				}
					c = append(c, v[0])
				}
			}
	}
	return
}

func (r *Commande) GetMessage(re *http.Request) (a []string) {
	re.ParseForm()
	if v, ok := re.Form["message"]; ok {
		return v
	}
	return nil
}

func (r *Commande) Think(re *http.Request) (a string) {
	re.ParseForm()
	if v, ok := re.Form["think"]; ok {
		if len(v) > 0 && v[0] == "on" {
			return "cowthink"
		}
	}
	return "cowsay"
}

func (r *Commande) MakeHaha(a string, re *http.Request) (b []Haha){
	re.ParseForm()
	tt := r.Co[a].Com
	for key, val := range re.Form{
			if len(val[0]) > 0 {
			for _, vals := range tt {
				if key == vals.A {
					for _, ee := range val {
						var gg Haha
						gg.B = key
						if vals.B > 0 {
							gg.C  = append(gg.C, ee)
						}
						b = append(b, gg)
					}
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
	f, err := os.OpenFile("linuxtool.log", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	var tt string = m.Com
	for _, e := range m.Opt {
		tt = tt + " " + e
	}
	tt = tt + "\n"
	if _, err := f.WriteString(tt); err != nil {
		log.Println(err)
	}
	cmd := exec.Command(m.Com, m.Opt...)
	cmd.Stdout = &out
	cmd.Stderr  = &er
	err = cmd.Start()
	if err != nil {
		return out, er, err
	}
	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(60 * time.Second):
		println("programe kill")
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println(err)
		}
		return
	case err := <-done:
		if err != nil {
			println(err.Error())
		}
	}
	return
}
