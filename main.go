package main

import (
	"fmt"
	"polaroid/admin"
	"polaroid/user"
	"polaroid/connect"
	"polaroid/config"
	"polaroid/server"
	"polaroid/lw"
)

func main() {
	conf := config.NewConf()
	conf["pk"] = config.Pk()
	//l := NewPk(conf)
	//db := l.Connect();
	srv := server.NewServer(conf)
	adm := admin.NewAdmin(srv.Data)
	srv.AddHH(adm)
	usr := user.NewUser(srv.Data)
	srv.AddHH(usr)
	llw := lw.NewLw(srv.Data)
	srv.AddHH(llw)
	connect := connect.NewConnect(srv.Data);
	srv.AddHH(connect);
	serve := srv.Servers(conf)
	err := serve.ListenAndServe()
	fmt.Println(err)
}
