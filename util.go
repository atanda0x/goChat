package main

import (
	"errors"
	"net/http"
)

// session fun helps check, if user is login or not with the help of cookie
func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {

	// retrieve the cookie from req
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
