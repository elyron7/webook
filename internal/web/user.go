package web

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) RegisterRouter(r *gin.Engine) {
	ug := r.Group("/user")

	ug.POST("/signup", u.Signup)   // Handles user signup
	ug.POST("/login", u.Login)     // Handles user login
	ug.POST("/edit", u.Edit)       // Handles user info editing
	ug.POST("/profile", u.Profile) // Handles user profile
}

func (u *UserHandler) Signup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "signup",
	})
}

func (u *UserHandler) Login(c *gin.Context) {}

func (u *UserHandler) Edit(c *gin.Context) {}

func (u *UserHandler) Profile(c *gin.Context) {}
