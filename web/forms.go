package web

import (
	"encoding/gob"
)

func init() {
	gob.Register(RegisterForm{})
	gob.Register(LoginForm{})
	gob.Register(FormErrors{})
}

type FormErrors map[string]string

type RegisterForm struct {
	Email      string
	Password   string
	EmailTaken bool

	Errors FormErrors
}

func (f *RegisterForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Email == "" {
		f.Errors["Email"] = "Please enter a email."
	} else if f.EmailTaken {
		f.Errors["Email"] = "This email is already taken."
	}

	if f.Password == "" {
		f.Errors["Password"] = "Please enter a password."
	} else if len(f.Password) < 8 {
		f.Errors["Password"] = "Your password must be at least 8 characters long."
	}

	return len(f.Errors) == 0
}

type LoginForm struct {
	Email                string
	Password             string
	IncorrectCredentials bool

	Errors FormErrors
}

func (f *LoginForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Email == "" {
		f.Errors["Email"] = "Please enter a email."
	} else if f.IncorrectCredentials {
		f.Errors["Email"] = "Email or password is incorrect."
	}

	if f.Password == "" {
		f.Errors["Password"] = "Please enter a password."
	}

	return len(f.Errors) == 0
}
