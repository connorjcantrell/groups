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

	user := UserHandler{store: store, sessions: sessions}
	group := GroupHandler{store: store, sessions: sessions}
	event := EventHandler{store: store, sessions: sessions}
	book := BookHandler{store: store, sessions: sessions}
	chapter := ChapterHandler{store: store, sessions: sessions}
	section := SectionHandler{store: store, sessions: sessions}

	h.Use(middleware.Logger)                         // log request body
	h.Use(csrf.Protect(csrfKey, csrf.Secure(false))) // CSRF is cross-site ... to protect user from fake websties
	h.Use(sessions.LoadAndSave)                      // Load and save user session via cookie/...
	h.Use(h.withUser)

	// RESTful APIs
	h.Post("/register", user.RegisterSubmit())

	h.Post("/login", user.LoginSubmit())

	h.Post("/groups", group.Store())
	h.Get("/groups/{id}", group.Group())

	h.Get("/groups/{id}/events") // 1-n relationship between group and events
	// group.GetEvents(),
	// event.GetByGroup(),
	// 2 functions are the same

	// CRU_ event -> 3 APIs -> miss 0
	// h.Get("/events/new", event.Create()) // HTML website
	h.Post("/events", event.Store()) // POST event -> use RESTful API convention
	// -> {
	//	"event": {"id": 1, "name": "my event"}
	// }
	h.Get("/events/{id}", event.Event())
	h.Delete("/events/{id}", event.Delete())

	h.Post("/books", book.Store())

	return h
}

type Handler struct {
	*chi.Mux

	userStore groups.UserStore
	sessions  *scs.SessionManager
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
