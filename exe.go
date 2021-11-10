package main

import (
	"fmt"
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
}

type Commande struct {
	Co map[string]Bash
}

func NewCommand() (r *Commande){
	r = &Commande{}
	r.Co = make(map[string]Bash)
	r.Co["fortune"] = MakeFortune()
	return
}

func (r *Commande) Make(a string, b []Haha, c []string) (m *Mimi, err error){
	if val, ok := r.Co[a]; ok {
		m = &Mimi{}
		m.Com = val.Bash
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

func main() {
	e := NewCommand()
	ee , _:= e.Make("fortune", nil, nil)
	u, uu, _ := e.Exec(ee)
	fmt.Println(u.String(), uu.String())
}
