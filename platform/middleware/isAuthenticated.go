// platform/middleware/isAuthenticated.go

package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticated(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.HTML(http.StatusOK, "home.html", nil)
	} else {
		ctx.Next()
	}
}
