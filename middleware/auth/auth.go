package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/auth"
	"github.com/valentinRog/sba-todo/handler"
	"github.com/valentinRog/sba-todo/store"
)

type Middleware struct {
	q   *store.Queries
	h   *handler.Handlers
	ctx context.Context
}

func New(ctx context.Context, h *handler.Handlers, q *store.Queries) *Middleware {
	return &Middleware{
		ctx: ctx,
		h:   h,
		q:   q,
	}
}

func (m *Middleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Path(), "/static") {
			return next(c)
		}

		cookie, _ := c.Request().Cookie("token")

		code := http.StatusMovedPermanently
		if c.Request().Header.Get("HX-Request") == "true" {
			code = 200
		}

		if cookie != nil {
			if _, ok := auth.GetUserIdFromToken(cookie.Value); !ok {
				c.SetCookie(auth.DeleteCookie())
				c.Response().Header().Set("HX-Redirect", "/login")
				return c.Redirect(code, "/login")
			}
		}

		fmt.Println(c.Path())

		if cookie == nil && !strings.HasPrefix(c.Path(), "/login") && !strings.HasPrefix(c.Path(), "/auth") {
			c.Response().Header().Set("HX-Redirect", "/login")
			return c.Redirect(code, "/login")
		}

		if cookie != nil && strings.HasPrefix(c.Path(), "/login") {
			return c.Redirect(http.StatusMovedPermanently, "/")
		}

		return next(c)
	}
}
