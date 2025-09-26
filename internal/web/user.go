package web

import (
	"errors"
	"net/http"

	"github.com/elyron7/webook/internal/domain"
	"github.com/elyron7/webook/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	ErrBadRequest        = errors.New("invalid json data")
	ErrSessionSaveFailed = errors.New("session save failed")
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) RegisterRouter(r *gin.Engine) {
	ug := r.Group("/users")

	ug.POST("/signup", u.Signup)   // Handles user signup
	ug.POST("/login", u.Login)     // Handles user login
	ug.POST("/edit", u.Edit)       // Handles user info editing
	ug.POST("/profile", u.Profile) // Handles user profile
}

func (u *UserHandler) Signup(c *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}

	var req SignUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrBadRequest.Error()})
		return
	}

	user := domain.User{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	if err := u.svc.SignUp(c.Request.Context(), user); err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, service.ErrInvalidEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successfully"})
}

func (u *UserHandler) Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrBadRequest.Error()})
		return
	}

	user := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := u.svc.Login(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("userId", user.Id)
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": ErrSessionSaveFailed.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}

func (u *UserHandler) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Edit"})
}

func (u *UserHandler) Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Profile"})
}
