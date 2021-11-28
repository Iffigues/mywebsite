package types

import (
	"polaroid/linuxtool"
	"polaroid/config"
	//"polaroid/pk"
	"github.com/gorilla/sessions"
)

type Data struct {
	Store *sessions.CookieStore
	Conf  config.Config
	Commande  *linuxtool.Commande
	//Db    *pk.Pk
}
