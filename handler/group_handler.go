package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/v2"
	db "github.com/connorjcantrell/groups/postgw/sqlc"
	"github.com/go-chi/chi"
)

type GroupHandler struct {
	sessions *scs.SessionManager
}

// higher-order function: a function that receives or returns a function
// higher-order function if you need more params than what you can pass (because the limitation of the library)

// you can use struct instead of higher-order function
func (h *GroupHandler) GetGroupByID(querier db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// https://osintiostom.medium.com/go-gorilla-mux-how-to-get-url-params-fa702901df38
		idString := chi.URLParam(r, "id")

		// string() -> creates a lot of problems. `string(123)` -> not "123", it will return a character from ASCII/UTF code

		id, err := strconv.Atoi(idString)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		group, err := querier.GetGroup(r.Context(), int32(id))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json") // set the content type to json
		data, err := json.Marshal(struct {
			Group db.Group `db:"group"`
		}{ //-> define a new struct
			Group: group,
		}) // -> use that new struct
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
		return
	}
}

// Store will store a new Group to the database using the form data provided by the user
func (h *GroupHandler) Store(querier db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		form := db.CreateGroupParams{
			Organizer:   1, // TODO: Get user id from scs.sessions
			Name:        r.FormValue("name"),
			Description: r.FormValue("Description"),
		}
		querier.CreateGroup(r.Context(), form)
	}
}
