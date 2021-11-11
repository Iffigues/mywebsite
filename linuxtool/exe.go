package linuxtool

import (
	"errors"
	"bytes"
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
	Co map[string]Bash
}

func NewCommand() (r *Commande){
	r = &Commande{}
	r.Co = make(map[string]Bash)
	r.Co["fortune"] = MakeFortune()
	r.Co["figlet"] = MakeFiglet()
	r.Co["toilet"] = MakeToilet()
	r.Co["cowsay"] = MakeCowsay()
	r.Co["cowthink"] = MakeCowthink()
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
