package web

import (
	"log"
	"net/http"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

const (
	emailRegex    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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

	ok, err := ValidateEmail(req.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "System error"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email format error"})
		return
	}

	ok, err = ValidatePassword(req.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "System error"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password format error"})
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (u *UserHandler) Login(c *gin.Context) {}

func (u *UserHandler) Edit(c *gin.Context) {}

func (u *UserHandler) Profile(c *gin.Context) {}

func ValidatePassword(password string) (bool, error) {
	re := regexp2.MustCompile(passwordRegex, 0)
	return re.MatchString(password)
}

func ValidateEmail(email string) (bool, error) {
	re := regexp2.MustCompile(emailRegex, 0)
	return re.MatchString(email)
}
