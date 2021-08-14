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

func NewHandler(sessions *scs.SessionManager, csrfKey []byte) *Handler {
	h := &Handler{
		Mux:      chi.NewMux(),
		store:    store,
		sessions: sessions,
	}

	// handler only needs user, not a big session. Maybe handlers can use use user directly
	user := sessions.GetUser()
	userHandler := UserHandler{user: user}
	groupHandler := GroupHandler{sessions: sessions}
	eventHandler := EventHandler{sessions: sessions}
	bookHandler := BookHandler{sessions: sessions}
	chapterHandler := ChapterHandler{sessions: sessions}
	sectionHandler := SectionHandler{sessions: sessions}

	h.Use(middleware.Logger)                         // log request body
	h.Use(csrf.Protect(csrfKey, csrf.Secure(false))) // CSRF is cross-site ... to protect user from fake websties
	h.Use(sessions.LoadAndSave)                      // Load and save user session via cookie/...
	h.Use(h.withUser)

	// RESTful APIs
	h.Post("/register", userHandler.RegisterSubmit())

	h.Post("/login", userHandler.LoginSubmit())

	h.Get("/groups", groupHandler.List())
	h.Get("/groups/{id}", groupHandler.Group())
	h.Post("/groups/new", groupHandler.Store())
	h.Post("/groups/{id}/update", groupHandler.Update())
	h.Post("groups/{id}/delete", groupHandler.Delete())

	// TODO: Implement 1-n relationship between group and events
	h.Get("/groups/{id}/events", groupHandler.GetEvents())

	h.Get("/events/{id}", eventHandler.Event())
	h.Post("/events/new", eventHandler.Store())
	h.Get("/events/user/{id}", eventHandler.EventsByUser())

	// CRU_ event -> 3 APIs -> miss 0
	// h.Get("/events/new", event.Create()) // HTML website

	// to check book is exist, use foreign key;
	// book_id 1 -> join table -> the DB will check whether the key is correct -> the book.id = 1 is exist
	// TODO: Read link
	// https://www.w3schools.com/sql/sql_foreignkey.asp

	// only do validations if you have all data for that in the application layer, if you need to query db, do not do it

	// 3 APIs:
	// - get groups: /groups
	// - get all books in a group: /groups/{id}/books
	// - get chapters in a book: /books/{id}/chapters (my preferred way) or /groups/{group_id}/books/{book_id}/chapters
	// -> get all chapters of a book in a group
	// - get all sections of a chapter: /chapters/{id}/sections?event={event_id}
	// 2 cases: -> how RESTful API works -> stick to the best practices and conventions
	// - if event query is exist: only get sections covered by the event
	// - if event is not exist: get all sections

	h.Post("/events", eventHandler.Store()) // POST event -> use RESTful API convention
	// -> {
	//	"event": {"id": 1, "name": "my event"}
	// }
	h.Get("/events/{id}", eventHandler.Event())
	h.Delete("/events/{id}", eventHandler.Delete())

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
