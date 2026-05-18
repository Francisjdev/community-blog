package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/database"
	"github.com/francisjdev/community-blog/internal/helper"
)

func RegisterUser(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		userPayload := createUserRequest{}
		err := decoder.Decode(&userPayload)
		if err != nil {
			helper.RespondWithError(w, 500, "Error decoding parameters")
			log.Printf("Error decoding parameters: %s", err)
			return
		}
		registerUserParams := database.CreateUserParams{
			Email:        userPayload.Email,
			PasswordHash: userPayload.Password,
		}
		returnedUserData, err := app.Service.Users.RegisterUser(r.Context(), registerUserParams)
		if err != nil {
			helper.RespondWithError(w, 500, "Error creating user")
			log.Printf("Error creating user: %s", err)
			return
		}
		user := User{
			Email: returnedUserData.Email,
		}
		helper.RespondWithJSON(w, 201, user)
		return
	}
}
