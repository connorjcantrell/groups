package form

// // Duplicate type declaration to keep errors in same scope as file (register.go and login.go)
// type NoEmailErr struct{}

// func (e NoEmailErr) Error() string {
// 	return "no email is given"
// }

// // Duplicate type declaration to keep errors in same scope as file (register.go and login.go)
// type NoPasswordErr struct{}

// func (e NoPasswordErr) Error() string {
// 	return "no password given"
// }

type LoginForm struct {
	Email    string
	Password string
}

func (f *LoginForm) Validate() (bool, error) {
	if f.Email == "" {
		return false, NoEmailErr{}
	}
	if f.Password == "" {
		return false, NoPasswordErr{}
	}
	return true, nil
}
