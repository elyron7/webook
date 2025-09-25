package web

import (
	"errors"
	"net/http"

	"github.com/elyron7/webook/internal/domain"
	"github.com/elyron7/webook/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	ErrEmailAlreadyExists = service.ErrEmailAlreadyExists
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "xxx"})
		return
	}

	user := domain.User{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	if err := u.svc.SignUp(c.Request.Context(), user); err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successfully"})
}

func (u *UserHandler) Login(c *gin.Context) {}

func (u *UserHandler) Edit(c *gin.Context) {}

func (u *UserHandler) Profile(c *gin.Context) {}
