package form

import "fmt"

type RegisterForm struct {
	Email    string
	Password string
}

// Errors
type EmailTakenErr struct{ email string }

func (e EmailTakenErr) Error() string {
	return fmt.Sprintf("the email %s is taken", e.email)
}

type NoEmailErr struct{}

func (e NoEmailErr) Error() string {
	return "no email is given"
}

type NoPasswordErr struct{}

func (e NoPasswordErr) Error() string {
	return "no password given"
}

type ShortPasswordErr struct{ passwordLen int }

func (e ShortPasswordErr) Error() string {
	return fmt.Sprintf("the password is not long enough, got %d, expected 8", e.passwordLen)
}

// Still needed by your API
func (f *RegisterForm) Validate() (bool, error) {

	if f.Email == "" {
		return false, NoEmailErr{}
	}
	if f.Password == "" {
		return false, NoPasswordErr{}
	} else if len(f.Password) < 8 {
		return false, ShortPasswordErr{len(f.Password)}
	}
	return true, nil
}

// func Caller() {
// 	result, err := Validate()
// 	if errors.Is(err, EmailTakenErr) {
// 		log.Error("taken error")
// 	} else if errors.Is(err, NoEmailErr) {
// 	}
// }

// loginForm.IncorrectCredentials
// loginForm.Validate()
