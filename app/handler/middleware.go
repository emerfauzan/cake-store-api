package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/emerfauzan/cake-store-api/lib"
	"github.com/emerfauzan/cake-store-api/lib/logger"
	"github.com/felixge/httpsnoop"
	"github.com/google/uuid"
)

func (handler *Handler) UserMiddlewares(next http.Handler) http.Handler {
	return handler.StandardMiddleware(handler.AuthMiddleware(next))
}

func (handler *Handler) PanicMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error(r.Context(), "panic occured", map[string]interface{}{
					"error": rec,
				})
				writeError(w, lib.ErrorInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (handler *Handler) StandardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "X-Request-ID", uuid.NewString())

		m := httpsnoop.CaptureMetrics(handler.PanicMiddlewares(next), w, r.WithContext(ctx))

		logger.Info(ctx, "http api request", map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL,
			"status":   m.Code,
			"duration": m.Duration.Milliseconds(),
		})
	})
}

func (handler *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")

		if len(splitToken) < 2 {
			writeError(w, lib.ErrorForbidden)
			return
		}

		accessToken := splitToken[1]
		session, err := handler.usecase.ValidateToken(r.Context(), accessToken)
		if err != nil {
			fmt.Println(err)
			writeError(w, lib.ErrorForbidden)
			return
		}

		if session.Id == 0 {
			writeError(w, lib.ErrorTokenNotValid)
			return
		}

		if session.ExpiresAt != nil && session.ExpiresAt.Before(time.Now()) {
			writeError(w, lib.ErrorSessionExpired)
			return
		}

		ctx := context.WithValue(r.Context(), "CurrentUserID", session.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
