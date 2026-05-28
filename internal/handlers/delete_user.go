package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/helper"
)

func DeleteUser(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		userPayload := User{}
		err := decoder.Decode(&userPayload)
		if err != nil {
			helper.RespondWithError(w, 500, "Error decoding parameters")
			log.Printf("Error decoding parameters: %s", err)
			return
		}
		err = app.Service.Users.DeleteUser(r.Context(), userPayload.Email)
		if err != nil {
			helper.RespondWithError(w, 500, "Error deleting user")
			log.Printf("Error deleting user: %s", err)
			return
		}

		helper.RespondWithJSON(w, 201, userPayload.Email)
		return
	}
}
