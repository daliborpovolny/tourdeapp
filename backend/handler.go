package main

import (
	"context"
	"net/http"
	"time"
	"tourbackend/crypto"
	db "tourbackend/database/gen"

	"github.com/labstack/echo/v4"
)

// The Handler is a struct which methods handle the individual endpoints
// the handler has a helper method called newReqCtx which adds creates a context for each request
// this context allows interaction with the database and has few helper methods like Error and JSONMsg
type Handler struct {
	queries *db.Queries
}

type RequestCtx struct {
	Ctx     context.Context
	Echo    echo.Context
	Queries *db.Queries

	User *db.User
}

func (h *Handler) newReqCtx(c echo.Context) *RequestCtx {

	r := &RequestCtx{
		Ctx:     c.Request().Context(),
		Echo:    c,
		Queries: h.queries,
	}

	if user, ok := c.Get("user").(*db.User); ok {
		r.User = user
	}

	return r
}

// Helpers

func (r *RequestCtx) Error(code int, msg string) error {
	return r.Echo.JSON(code, map[string]string{"error": msg})
}

func (r *RequestCtx) JSONMsg(code int, msg string) error {
	return r.Echo.JSON(code, map[string]string{"message": msg})
}

// Auth stuff

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) login(c echo.Context) error {
	r := h.newReqCtx(c)

	var req LoginRequest
	if err := r.Echo.Bind(&req); err != nil {
		return r.Error(http.StatusBadRequest, "invalid request body")
	}

	user, err := r.Queries.GetUserByEmail(r.Ctx, req.Email)
	if err != nil {
		return err
	}

	isCorrect := crypto.CheckPasswordHash(req.Password, user.Hash)
	if !isCorrect {
		return r.Error(http.StatusUnauthorized, "invalid password")
	}

	newToken, err := crypto.NewSessionToken()
	if err != nil {
		c.Logger().Errorf("failed to generate a session token: %v", err)
		return r.Error(http.StatusInternalServerError, "internal server error")
	}

	cookie := createHttpCookie(newToken)

	_, err = r.Queries.CreateSession(r.Ctx, db.CreateSessionParams{
		UserID:    user.ID,
		Token:     newToken,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: cookie.Expires.Unix(),
	})
	if err != nil {
		c.Logger().Errorf("failed to create the session in the database: %v", err)
		return r.Error(http.StatusInternalServerError, "internal server error")
	}

	c.SetCookie(cookie)

	c.Logger().Infof("logged in a user: %v", user.Email)
	return r.JSONMsg(http.StatusCreated, "logged in user")
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (h *Handler) register(c echo.Context) error {
	r := h.newReqCtx(c)

	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return r.Error(http.StatusBadRequest, "invalid request")
	}

	hash, err := crypto.HashPassword(req.Password)
	if err != nil {
		return r.Error(http.StatusBadRequest, "unhashable password")
	}

	user, err := r.Queries.CreateUser(r.Ctx, db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Hash:      hash,
		Email:     req.Email,
	})
	if err != nil {
		if IsUniqueConstraintError(err) {
			return r.Error(http.StatusBadRequest, "email already in use")
		}

		c.Logger().Errorf("failed to create a user: %v", err)
		return r.Error(http.StatusInternalServerError, "internal server error")
	}

	newToken, err := crypto.NewSessionToken()
	if err != nil {
		c.Logger().Errorf("failed to generate a session token: %v", err)
		return r.Error(http.StatusInternalServerError, "internal server error")
	}

	cookie := createHttpCookie(newToken)

	_, err = r.Queries.CreateSession(r.Ctx, db.CreateSessionParams{
		UserID:    user.ID,
		Token:     newToken,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: cookie.Expires.Unix(),
	})
	if err != nil {
		c.Logger().Errorf("failed to create the session in the database: %v", err)
		return r.Error(http.StatusInternalServerError, "internal server error")
	}

	c.SetCookie(cookie)

	c.Logger().Infof("registered a user: %v", user.Email)
	return r.JSONMsg(http.StatusCreated, "registered user")
}

func (h *Handler) profile(c echo.Context) error {
	r := h.newReqCtx(c)

	if r.User == nil {
		return r.Error(http.StatusUnauthorized, "authentication required")
	}

	c.Logger().Infof("returned profile of user: %v", r.User.Email)

	return c.JSON(http.StatusOK, map[string]string{
		"first_name": r.User.FirstName,
		"last_name":  r.User.LastName,
		"email":      r.User.Email,
	})
}
