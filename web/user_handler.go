package web

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"
	groups "github.com/connorjcantrell/groups"
	"github.com/connorjcantrell/groups/models"
	"github.com/connorjcantrell/groups/web/form"
	"github.com/connorjcantrell/groups/postgres"
	"&
	github.com/google/uuid"
	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	store    models.User
	sessions *scs.SessionManager
}

func (h *UserHandler) Register() http.HandlerFunc {
	type data struct {
		SessionData
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/user_register.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data{
			SessionData: GetSessionData(h.sessions, r.Context()),
		})
	}
}

func (h *UserHandler) RegisterSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := form.RegisterForm{
			Email:    r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		if _, err := postgres..UserByEmail(form.Email); err == nil {
			form.EmailTaken = true
		}
		b, err := form.Validate()
		if !b {
			h.sessions.Put(r.Context(), "form", form)
			http.Redirect(w, r, r.Referer(), http.StatusFound)
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := h.store.CreateUser(&groups.User{
			ID:       uuid.New(),
			Email:    form.Email,
			Password: string(password),
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.sessions.Put(r.Context(), "flash", "Your registration was successful. Please log in.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *UserHandler) LoginSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := LoginForm{
			Email:                r.FormValue("username"),
			Password:             r.FormValue("password"),
			IncorrectCredentials: false,
		}
		user, err := h.UserByEmail(form.Email)
		if err != nil {
			form.IncorrectCredentials = true
		} else {
			compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
			form.IncorrectCredentials = compareErr != nil
		}
		if !form.Validate() {
			h.sessions.Put(r.Context(), "form", form)
			http.Redirect(w, r, r.Referer(), http.StatusFound)
			return
		}

		h.sessions.Put(r.Context(), "user_id", user.ID)
		h.sessions.Put(r.Context(), "flash", "You have been logged in sucessfully.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *UserHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.sessions.Remove(r.Context(), "user_id")
		h.sessions.Put(r.Context(), "flash", "You have been logged out sucessfully.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
