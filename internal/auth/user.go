package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	Register(c *gin.Context, user *User) error
	Login(c *gin.Context, email, password string) (string, error)
	GetUser(c *gin.Context, userID uint) (*User, error)
}

type userSvc struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userSvc{db: db}
}

func (u *userSvc) Register(c *gin.Context, user *User) error {
	// Implement user registration logic here
	// Example code to insert user record into database
	result := u.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userSvc) Login(c *gin.Context, email, password string) (string, error) {
	// Implement user login logic here
	// Example code to authenticate user credentials and return JWT token
	// ...
	return "jwt_token", nil
}

func (u *userSvc) GetUser(c *gin.Context, userID uint) (*User, error) {
	// Implement code to retrieve user from database by ID
	var user User
	result := u.db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UserAPIHandler(userSvc UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost:
			var user User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := userSvc.Register(c, &user); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
		case http.MethodGet:
			userID, err := getUserIDFromContext(c)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			user, err := userSvc.GetUser(c, userID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

func getUserIDFromContext(c *gin.Context) (uint, error) {
	userID, ok := c.Get("userID")
	if !ok {
		return 0, ErrUnauthorized
	}
	return userID.(uint), nil
}
