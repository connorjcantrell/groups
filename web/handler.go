package web

import (
	"context"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/connorjcantrell/groups"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

func NewHandler(store groups.Store, sessions *scs.SessionManager, csrfKey []byte) *Handler {
	h := &Handler{
		Mux:      chi.NewMux(),
		store:    store,
		sessions: sessions,
	}

	users := UserHandler{store: store, sessions: sessions}
	groups := GroupHandler{store: store, sessions: sessions}
	events := EventHandler{store: store, sessions: sessions}
	books := BookHandler{store: store, sessions: sessions}
	chapters := ChapterHandler{store: store, sessions: sessions}
	sections := SectionHandler{store: store, sessions: sessions}

	h.Use(middleware.Logger)
	h.Use(csrf.Protect(csrfKey, csrf.Secure(false)))
	h.Use(sessions.LoadAndSave)
	h.Use(h.withUser)

	return h
}

type Handler struct {
	*chi.Mux

	store    groups.Store
	sessions *scs.SessionManager
}

func (h *Handler) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := h.sessions.Get(r.Context(), "user_id").(uuid.UUID)

		user, err := h.store.User(id)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
