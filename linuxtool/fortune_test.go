package linuxtool

import (
	"testing"
	"fmt"
)

func TestSum1(t *testing.T) {
	e := NewCommand()
	ee , err:= e.Make("fortune", []Haha{
		Haha{"-n", []string{}},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ee)
	u, uu, et := e.Exec(ee)
	if et != nil {
		t.Fatal(et)
	}
	t.Fatal(u.String(), uu.String())
}

func TestSum(t *testing.T) {
	e := NewCommand()
	ee , err:= e.Make("fortune", []Haha{
		Haha{"-n", []string{"er"}},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ee)
	u, uu, et := e.Exec(ee)
	if et != nil {
		t.Fatal(et)
	}
	t.Fatal(u.String(), uu.String())
}
