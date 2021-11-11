package user

import (
	"net/http"
	"polaroid/server"
	"polaroid/types"
)

type User struct {
	Data *types.Data
}

type NewUsers struct {
	Email string `json:"email"`
	Login string `json:"login"`
	Pwd string `json:"pwd"`
	Pwd1 string `json:"pwd1"`

}

func NewUser(s *types.Data) (a *User) {
	a = new(User)
	a.Data = s
	return
}

func (a *User) Signup(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *User) Loging(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}


func (a *User) Logout(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *User) WWW(s *server.Server) {
	s.NewR("/signup", "signup", []string{"GET","POST"}, a.Signup(s.Data), 1)
	s.NewR("/signin", "signin", []string{"GET", "POST"}, a.Loging(s.Data), 1)
	s.NewR("/signout", "signout", []string{"GET"}, a.Logout(s.Data), 2)
}
