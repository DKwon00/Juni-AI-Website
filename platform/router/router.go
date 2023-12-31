// platform/router/router.go

package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"01-Login/platform/authenticator"
	"01-Login/platform/middleware"
	"01-Login/web/app/callback"
	"01-Login/web/app/login"
	"01-Login/web/app/logout"
	"01-Login/web/app/user"
	"01-Login/web/app/userdb"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		if sessions.Default(ctx).Get("profile") == nil {
			ctx.HTML(http.StatusOK, "home.html", nil)
		} else {
			session := sessions.Default(ctx)
			profile := session.Get("profile")

			ctx.HTML(http.StatusOK, "user.html", profile)
		}

	})

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	//known bug, will load user.html and home.html when visiting /user directly while signed out
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)
	router.GET("/userdb", userdb.Handler)

	return router
}
