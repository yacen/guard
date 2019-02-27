package handler

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/yacen/guard/service"

	"github.com/yacen/guard/config"
	"github.com/yacen/guard/util/log"

	"github.com/yacen/guard/resp"
	"github.com/yacen/guard/util"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Error(errors.New("post method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Error(err)
		util.WriteMessage(w, resp.ParseFrom, "parse form error")
		return
	}

	username := r.PostFormValue("username")
	ok, err := regexp.MatchString(config.Cfg.UsernamePattern, username)
	if err != nil {
		log.Error(err)
		util.WriteMessage(w, resp.ParseFrom, "match username error")
		return
	}
	if !ok {
		log.Println("check username error")
		util.WriteMessage(w, resp.ParseFrom, "check username error")
		return
	}

	password := r.PostFormValue("password")
	ok, err = regexp.MatchString(config.Cfg.PasswordPattern, password)
	if err != nil {
		log.Error(err)
		util.WriteMessage(w, resp.ParseFrom, "match password error")
		return
	}
	if !ok {
		util.WriteMessage(w, resp.ParseFrom, "check password error")
		return
	}

	token, err := service.SignIn(username, password)
	if err != nil {
		log.Error(err)
		util.WriteMessage(w, resp.ParseFrom, "sign in failed")
		return
	}
	util.WriteData(w, resp.Token{Token: string(token)})
}
