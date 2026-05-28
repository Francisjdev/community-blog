package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/helper"
)

func Login(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		userPayload := createUserRequest{}
		err := decoder.Decode(&userPayload)
		if err != nil {
			helper.RespondWithError(w, 500, "Error decoding parameters")
			log.Printf("Error decoding parameters: %s", err)
			return
		}

		accessToken, refreshToken, err := app.Service.Users.Login(r.Context(), userPayload.Email, userPayload.Password, app.Config.SecretKey)
		if err != nil {
			helper.RespondWithError(w, 400, "error")
			log.Printf("Error login user: %s", err)
			return
		}
		type blabli struct {
			AToken string `json:"access_token"`
			RToken string `json:"refresh_token"`
		}
		response := blabli{
			AToken: accessToken,
			RToken: refreshToken,
		}
		helper.RespondWithJSON(w, 200, response)
	}
}
