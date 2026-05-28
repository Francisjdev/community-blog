package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/auth"
	"github.com/francisjdev/community-blog/internal/helper"
)

type contextKey string

const ClaimsKey contextKey = "claims"

func AuthMiddleware(app *app.Application, next http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		extractedToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			helper.RespondWithError(w, 401, "Not authorized")
			log.Printf("error: %v\n", err)
			return
		}
		validatedTokenClaims, err := auth.ValidateJWt(extractedToken, app.Config.SecretKey)

		if err != nil {
			helper.RespondWithError(w, 401, "Not authorized")
			log.Printf("error: %v\n", err)
			return
			//custom error
		}
		ctx := context.WithValue(r.Context(), ClaimsKey, validatedTokenClaims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
