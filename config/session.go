package config

import "github.com/gorilla/sessions"

const SESSION_ID = "go_auth_sess"

var Store = sessions.NewCookieStore([]byte("ertyuioasdkjasdhiuq"))
