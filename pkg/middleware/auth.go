package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rows15/CTDGolangBackendIII/pkg/web"
)

func Authentication() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	return func(ctx *gin.Context) {
		tokenReceived := ctx.GetHeader("SECRET_TOKEN")

		if tokenReceived == "" {
			web.BadResponse(ctx, http.StatusUnauthorized, "error", "Token not found")
			ctx.Abort()
			return
		}

		if tokenReceived != requiredToken {
			web.BadResponse(ctx, http.StatusUnauthorized, "error", "Invalid token provided")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
